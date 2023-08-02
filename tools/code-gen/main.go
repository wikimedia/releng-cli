package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dave/jennifer/jen"
	"github.com/sirupsen/logrus"
	"gitlab.wikimedia.org/repos/releng/cli/pkg/codegen"
	"gopkg.in/yaml.v2"
)

type Spec []Command

type Command struct {
	Use         string    `yaml:"use"`
	Aliases     []string  `yaml:"aliases,omitempty"`
	Short       string    `yaml:"short,omitempty"`
	Example     string    `yaml:"example,omitempty"`
	StringFlags []Flag    `yaml:"string-flags,omitempty"`
	SubCommands []Command `yaml:"sub-commands,omitempty"`
	GerritPath  string    `yaml:"gerrit-path,omitempty"`
	HttpMethod  string    `yaml:"http-method,omitempty"`
}

type Flag struct {
	Name        string `yaml:"name,omitempty"`
	Required    bool   `yaml:"required,omitempty"`
	Value       string `yaml:"value,omitempty"`
	Usage       string `yaml:"usage,omitempty"`
	GerritParam string `yaml:"gerrit-param,omitempty"`
	Body        bool   `yaml:"body,omitempty"`
}

func main() {
	// Load YAMl file
	yamlFile, err := ioutil.ReadFile("tools/code-gen/gerrit.yml")
	if err != nil {
		logrus.Fatal(err)
	}
	var spec Spec
	err = yaml.Unmarshal(yamlFile, &spec)
	if err != nil {
		logrus.Fatal(err)
	}

	reGenerateGerritCommands(spec)
}

func deleteGeneratedFilesFromDirectory(dir string) {
	// Delete all files that end in .gen.go
	files, err := filepath.Glob(dir + "/*.gen.go")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}

func reGenerateGerritCommands(spec Spec) {
	deleteGeneratedFilesFromDirectory("internal/cmd/gerrit")
	for _, c := range spec {
		f := jen.NewFile("gerrit")
		f.Comment("This code is generated by tools/code-gen/main.go. DO NOT EDIT.")

		f.Add(cobraCommandFuncWithSubCommands(c, "")...)

		filePath := "internal/cmd/gerrit/" + codegen.ForFileName(c.Use) + ".gen.go"
		err := f.Save(filePath)
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Println(filePath)
	}
}

func cobraCommandFuncWithSubCommands(c Command, parentName string) []jen.Code {
	var funcs []jen.Code
	currentLevelname := codegen.ForFunctionNamePart(parentName) + codegen.ForFunctionNamePart(c.Use)
	funcs = append(funcs, cobraCommandFunc(currentLevelname, c))
	for _, subC := range c.SubCommands {
		funcs = append(funcs, cobraCommandFuncWithSubCommands(subC, currentLevelname)...)
	}
	return funcs
}

func cobraCommandFunc(namePartForFunc string, c Command) jen.Code {
	return jen.Func().Id("NewGerrit" + namePartForFunc + "Cmd").Params().Op("*").Op("cobra.Command").Add(cobraCommandBlock(namePartForFunc, c)).Line()
}

func cobraCommandBlock(namePartForFunc string, c Command) jen.Code {
	var addCommandsToCmd []jen.Code
	for _, subC := range c.SubCommands {
		addCommandsToCmd = append(addCommandsToCmd, jen.Id("cmd").Dot("AddCommand").Call(jen.Id("NewGerrit"+namePartForFunc+codegen.ForFunctionNamePart(subC.Use)+"Cmd").Call()).Line())
	}

	var addFlagsToCmd []jen.Code
	var defineFlags []jen.Code
	for _, flag := range c.StringFlags {
		addFlagsToCmd = append(addFlagsToCmd, jen.Id("cmd").Dot("Flags").Call().Dot("StringVar").Call(
			jen.Op("&").Id("cmdFlags").Dot(flag.Name),
			jen.Lit(flag.Name),
			jen.Lit(flag.Value),
			jen.Lit(flag.Usage),
		).Line())
		if flag.Required {
			addFlagsToCmd = append(addFlagsToCmd, jen.Id("cmd").Dot("MarkFlagRequired").Call(jen.Lit(flag.Name)).Line())
		}
		defineFlags = append(defineFlags, jen.Id(flag.Name).String())
	}

	blockParts := []jen.Code{}
	if c.StringFlags != nil {
		blockParts = append(blockParts, jen.Add(jen.Type().Id("flags").Struct(defineFlags...)).Line())
		blockParts = append(blockParts, jen.Id("cmdFlags").Op(":=").Id("flags").Block().Line())
	}
	blockParts = append(blockParts,
		jen.Id("cmd").Op(":=").Add(cobraCommandDefinition(c)).Line(),
		jen.Add(addCommandsToCmd...),
		jen.Add(addFlagsToCmd...),
		jen.Return(jen.Id("cmd")),
	)

	return jen.Block(jen.Add(blockParts...))
}

func cobraCommandDefinition(c Command) jen.Code {
	run := jen.Func().Params(jen.Id("cmd").Op("*").Qual("github.com/spf13/cobra", "Command"), jen.Id("args").Index().String()).Block()

	if c.GerritPath != "" {
		body := jen.Nil()

		pathReplacementSteps := []jen.Code{}
		for _, flag := range c.StringFlags {
			if flag.Body {
				body = jen.Id("cmdFlags").Dot(flag.Name)
				continue
			}
			lookFor := flag.Name
			if flag.GerritParam != "" {
				lookFor = flag.GerritParam
			}
			pathReplacementSteps = append(pathReplacementSteps, jen.Id("path").Op("=").Id("addParamToPath").Call(jen.Id("path"), jen.Lit(lookFor), jen.Id("cmdFlags").Dot(flag.Name)).Line())
		}

		httpMethod := "GET"
		if c.HttpMethod != "" {
			httpMethod = c.HttpMethod
		}

		run = jen.Func().Params(jen.Id("cmd").Op("*").Qual("github.com/spf13/cobra", "Command"), jen.Id("args").Index().String()).Block(
			// Define the URL path
			jen.Id("path").Op(":=").Lit(c.GerritPath),
			jen.Add(pathReplacementSteps...),

			jen.Id("client").Op(":=").Id("authenticatedClient").Call(),

			// Do the query & handle response
			jen.List(jen.Id("response"), jen.Id("err")).Op(":=").Id("client").Dot("Call").Call(jen.Lit(httpMethod), jen.Id("path"), body, jen.Nil()),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Qual("github.com/sirupsen/logrus", "Error").Call(jen.Id("err")),
			),
			jen.Defer().Id("response").Dot("Body").Dot("Close").Call(),
			jen.List(jen.Id("body"), jen.Id("err")).Op(":=").Qual("io/ioutil", "ReadAll").Call(jen.Id("response").Dot("Body")),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Panic(jen.Id("err")),
			),
			jen.Id("body").Op("=").Qual("github.com/andygrunwald/go-gerrit", "RemoveMagicPrefixLine").Call(jen.Id("body")),
			jen.Qual("fmt", "Print").Call(jen.Id("string").Call(jen.Id("body"))),
		)
	}

	cmdDict := jen.Dict{
		jen.Id("Use"):     jen.Lit(c.Use),
		jen.Id("Short"):   jen.Lit(c.Short),
		jen.Id("Example"): jen.Lit(c.Example),
	}
	if c.Aliases != nil {
		cmdDict[jen.Id("Aliases")] = jen.Index().String().ValuesFunc(func(g *jen.Group) {
			for _, alias := range c.Aliases {
				g.Lit(alias)
			}
		})
	}
	if c.GerritPath != "" {
		cmdDict[jen.Id("Run")] = run
	}

	return jen.Op("&").Qual("github.com/spf13/cobra", "Command").Block(cmdDict)
}
