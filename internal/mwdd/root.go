package mwdd

import (
	"os"

	"github.com/sirupsen/logrus"
	"gitlab.wikimedia.org/repos/releng/cli/internal/cli"
	"gitlab.wikimedia.org/repos/releng/cli/internal/mwdd/files"
)

/*MWDD representation of a mwdd v2 setup.*/
type MWDD struct {
	BaseDir string
}

var Context string

/*DefaultForUser returns the default mwdd working directory for the user.*/
func DefaultForUser() MWDD {
	return MWDD{
		BaseDir: mwddUserDirectory(),
	}
}

func mwddUserDirectory() string {
	return cli.UserDirectoryPathForCmd("mwdd")
}

/*Directory the directory containing the development environment.*/
func (m MWDD) Directory() string {
	return m.BaseDir + string(os.PathSeparator) + Context
}

/*EnsureReady ...*/
func (m MWDD) EnsureReady() {
	files.EnsureReady(m.Directory())
	m.Env().EnsureExists()
}

func (m MWDD) ComposerLocalJsonPath() string {
	// TODO move to internal.mediawiki
	return m.Env().Get("MEDIAWIKI_VOLUMES_CODE") + string(os.PathSeparator) + "composer.local.json"
}

func (m MWDD) ComposerLocalJsonExists() bool {
	// TODO move to internal.mediawiki
	_, err := os.Stat(m.ComposerLocalJsonPath())
	return !os.IsNotExist(err)
}

func (m MWDD) LocalSettingsPath() string {
	// TODO move to internal.mediawiki
	return m.Env().Get("MEDIAWIKI_VOLUMES_CODE") + string(os.PathSeparator) + "LocalSettings.php"
}

func (m MWDD) LocalSettingsContents() string {
	// TODO move to internal.mediawiki
	bytes, err := os.ReadFile(m.LocalSettingsPath())
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
	return string(bytes)
}

func (m MWDD) ExtensionsCheckedOut() []string {
	// TODO move to internal.mediawiki
	return directoriesInDirectory(m.Env().Get("MEDIAWIKI_VOLUMES_CODE") + string(os.PathSeparator) + "extensions")
}

func (m MWDD) SkinsCheckedOut() []string {
	// TODO move to internal.mediawiki
	return directoriesInDirectory(m.Env().Get("MEDIAWIKI_VOLUMES_CODE") + string(os.PathSeparator) + "skins")
}

func directoriesInDirectory(directory string) []string {
	entries, err := os.ReadDir(directory)
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
	directories := []string{}
	for _, e := range entries {
		if e.IsDir() {
			directories = append(directories, e.Name())
		}
	}
	return directories
}
