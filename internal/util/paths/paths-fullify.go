package paths

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

/*FullifyUserProvidedPath fullify people entering ~/ or ./ paths and them not being handeled anywhere.*/
func FullifyUserProvidedPath(userProvidedPath string) string {
	currentWorkingDirectory, _ := os.Getwd()

	if userProvidedPath == "." {
		return currentWorkingDirectory
	}
	if strings.HasPrefix(userProvidedPath, "./") {
		return filepath.Join(currentWorkingDirectory, userProvidedPath[2:])
	}

	usr, _ := user.Current()
	usrDir := usr.HomeDir

	if userProvidedPath == "~" {
		return usrDir
	}
	if strings.HasPrefix(userProvidedPath, "~/") {
		return filepath.Join(usrDir, userProvidedPath[2:])
	}

	// Fallback to what we were provided
	return userProvidedPath
}
