package gerrit

import (
	"fmt"
	gogerrit "github.com/andygrunwald/go-gerrit"
	logrus "github.com/sirupsen/logrus"
	cobra "github.com/spf13/cobra"
	"io"
)

// This code is generated by tools/code-gen/main.go. DO NOT EDIT.
func NewGerritProjectsCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Project Endpoints",
		Use:     "projects",
	}
	cmd.AddCommand(NewGerritProjectsListCmd())
	cmd.AddCommand(NewGerritProjectsGetCmd())
	cmd.AddCommand(NewGerritProjectsDescriptionCmd())
	cmd.AddCommand(NewGerritProjectsParentCmd())
	cmd.AddCommand(NewGerritProjectsHeadCmd())
	cmd.AddCommand(NewGerritProjectsConfigCmd())
	cmd.AddCommand(NewGerritProjectsAccessCmd())
	cmd.AddCommand(NewGerritProjectsBranchesCmd())
	cmd.AddCommand(NewGerritProjectsChildrenCmd())
	cmd.AddCommand(NewGerritProjectsTagsCmd())
	cmd.AddCommand(NewGerritProjectsDashboardsCmd())
	cmd.AddCommand(NewGerritProjectsLabelsCmd())
	cmd.AddCommand(NewGerritProjectsSubmitRequirementsCmd())
	return cmd
}
func NewGerritProjectsListCmd() *cobra.Command {
	type flags struct {
		query string
		limit string
		start string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/"
			path = addParamToPath(path, "query", cmdFlags.query)
			path = addParamToPath(path, "limit", cmdFlags.limit)
			path = addParamToPath(path, "start", cmdFlags.start)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "List Projects",
		Use:   "list",
	}
	cmd.Flags().StringVar(&cmdFlags.query, "query", "", "The query string to use to find projects.")
	cmd.Flags().StringVar(&cmdFlags.limit, "limit", "", "The maximum number of records to return.")
	cmd.Flags().StringVar(&cmdFlags.start, "start", "", "The index of the first record to return.")
	return cmd
}
func NewGerritProjectsGetCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Get a Project",
		Use:   "get",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsDescriptionCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Project description",
		Use:     "description",
	}
	cmd.AddCommand(NewGerritProjectsDescriptionGetCmd())
	return cmd
}
func NewGerritProjectsDescriptionGetCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/description/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Retrieves the description of a project.",
		Use:   "get",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsParentCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Project description",
		Use:     "parent",
	}
	cmd.AddCommand(NewGerritProjectsParentGetCmd())
	return cmd
}
func NewGerritProjectsParentGetCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/parent/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Retrieves the parent of a project.",
		Use:   "get",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsHeadCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Project HEAD",
		Use:     "head",
	}
	cmd.AddCommand(NewGerritProjectsHeadGetCmd())
	return cmd
}
func NewGerritProjectsHeadGetCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/HEAD/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Retrieves the HEAD of a project.",
		Use:   "get",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsConfigCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Get a Project config",
		Use:     "config",
	}
	cmd.AddCommand(NewGerritProjectsConfigGetCmd())
	return cmd
}
func NewGerritProjectsConfigGetCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/config/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Retrieves the config of a project.",
		Use:   "get",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsAccessCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Get a Project access",
		Use:     "access",
	}
	cmd.AddCommand(NewGerritProjectsAccessListCmd())
	return cmd
}
func NewGerritProjectsAccessListCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/access/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Lists the access of a project.",
		Use:   "list",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsBranchesCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Get a Project branches",
		Use:     "branches",
	}
	cmd.AddCommand(NewGerritProjectsBranchesListCmd())
	cmd.AddCommand(NewGerritProjectsBranchesGetCmd())
	return cmd
}
func NewGerritProjectsBranchesListCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/branches/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Lists the branches of a project.",
		Use:   "list",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsBranchesGetCmd() *cobra.Command {
	type flags struct {
		project string
		branch  string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/branches/{branch-name}/"
			path = addParamToPath(path, "project-name", cmdFlags.project)
			path = addParamToPath(path, "branch-name", cmdFlags.branch)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Retrieves a branch of a project.",
		Use:   "get",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	cmd.Flags().StringVar(&cmdFlags.branch, "branch", "", "The branch to retrieve.")
	cmd.MarkFlagRequired("branch")
	return cmd
}
func NewGerritProjectsChildrenCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Get a Project children",
		Use:     "children",
	}
	cmd.AddCommand(NewGerritProjectsChildrenListCmd())
	cmd.AddCommand(NewGerritProjectsChildrenGetCmd())
	return cmd
}
func NewGerritProjectsChildrenListCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/children/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Lists the children of a project.",
		Use:   "list",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsChildrenGetCmd() *cobra.Command {
	type flags struct {
		project string
		child   string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/children/{child-name}/"
			path = addParamToPath(path, "project-name", cmdFlags.project)
			path = addParamToPath(path, "child-name", cmdFlags.child)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Retrieves a child of a project.",
		Use:   "get",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	cmd.Flags().StringVar(&cmdFlags.child, "child", "", "The child to retrieve.")
	cmd.MarkFlagRequired("child")
	return cmd
}
func NewGerritProjectsTagsCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Get a Project tags",
		Use:     "tags",
	}
	cmd.AddCommand(NewGerritProjectsTagsListCmd())
	cmd.AddCommand(NewGerritProjectsTagsGetCmd())
	return cmd
}
func NewGerritProjectsTagsListCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/tags/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Lists the tags of a project.",
		Use:   "list",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsTagsGetCmd() *cobra.Command {
	type flags struct {
		project string
		tag     string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/tags/{tag-name}/"
			path = addParamToPath(path, "project-name", cmdFlags.project)
			path = addParamToPath(path, "tag-name", cmdFlags.tag)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Retrieves a tag of a project.",
		Use:   "get",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	cmd.Flags().StringVar(&cmdFlags.tag, "tag", "", "The tag to retrieve.")
	cmd.MarkFlagRequired("tag")
	return cmd
}
func NewGerritProjectsDashboardsCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Get a Project dashboards",
		Use:     "dashboards",
	}
	cmd.AddCommand(NewGerritProjectsDashboardsListCmd())
	cmd.AddCommand(NewGerritProjectsDashboardsGetCmd())
	return cmd
}
func NewGerritProjectsDashboardsListCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/dashboards/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Lists the dashboards of a project.",
		Use:   "list",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsDashboardsGetCmd() *cobra.Command {
	type flags struct {
		project   string
		dashboard string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/dashboards/{dashboard-name}/"
			path = addParamToPath(path, "project-name", cmdFlags.project)
			path = addParamToPath(path, "dashboard-name", cmdFlags.dashboard)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Retrieves a dashboard of a project.",
		Use:   "get",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	cmd.Flags().StringVar(&cmdFlags.dashboard, "dashboard", "", "The dashboard to retrieve.")
	cmd.MarkFlagRequired("dashboard")
	return cmd
}
func NewGerritProjectsLabelsCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Get a Project labels",
		Use:     "labels",
	}
	cmd.AddCommand(NewGerritProjectsLabelsListCmd())
	cmd.AddCommand(NewGerritProjectsLabelsGetCmd())
	return cmd
}
func NewGerritProjectsLabelsListCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/labels/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Lists the labels of a project.",
		Use:   "list",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsLabelsGetCmd() *cobra.Command {
	type flags struct {
		project string
		label   string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/labels/{label-name}/"
			path = addParamToPath(path, "project-name", cmdFlags.project)
			path = addParamToPath(path, "label-name", cmdFlags.label)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Retrieves a label of a project.",
		Use:   "get",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	cmd.Flags().StringVar(&cmdFlags.label, "label", "", "The label to retrieve.")
	cmd.MarkFlagRequired("label")
	return cmd
}
func NewGerritProjectsSubmitRequirementsCmd() *cobra.Command {
	cmd := &cobra.Command{

		Example: "",
		Short:   "Get a Project submit_requirements",
		Use:     "submit_requirements",
	}
	cmd.AddCommand(NewGerritProjectsSubmitRequirementsListCmd())
	cmd.AddCommand(NewGerritProjectsSubmitRequirementsGetCmd())
	return cmd
}
func NewGerritProjectsSubmitRequirementsListCmd() *cobra.Command {
	type flags struct {
		project string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/submit_requirements/"
			path = addParamToPath(path, "project-name", cmdFlags.project)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Lists the submit_requirements of a project.",
		Use:   "list",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	return cmd
}
func NewGerritProjectsSubmitRequirementsGetCmd() *cobra.Command {
	type flags struct {
		project            string
		submit_requirement string
	}
	cmdFlags := flags{}
	cmd := &cobra.Command{

		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/projects/{project-name}/submit_requirements/{submit_requirement-name}/"
			path = addParamToPath(path, "project-name", cmdFlags.project)
			path = addParamToPath(path, "submit_requirement-name", cmdFlags.submit_requirement)

			client := authenticatedClient(cmd.Context())
			response, err := client.Call(cmd.Context(), "GET", path, nil, nil)
			if err != nil {
				logrus.Error(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			body = gogerrit.RemoveMagicPrefixLine(body)
			fmt.Print(string(body))
		},
		Short: "Retrieves a submit_requirement of a project.",
		Use:   "get",
	}
	cmd.Flags().StringVar(&cmdFlags.project, "project", "", "The project to retrieve.")
	cmd.MarkFlagRequired("project")
	cmd.Flags().StringVar(&cmdFlags.submit_requirement, "submit_requirement", "", "The submit_requirement to retrieve.")
	cmd.MarkFlagRequired("submit_requirement")
	return cmd
}
