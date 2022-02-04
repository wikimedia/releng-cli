package cli

import (
	"os"

	"gitlab.wikimedia.org/releng/cli/internal/util/dirs"
)

// MWCLIDIR name of the directory for storing application files
const MWCLIDIR string = ".mwcli"

func Context() string {
	_, inGitlabCi := os.LookupEnv("GITLAB_CI")
	if !inGitlabCi && os.Getenv("MWCLI_CONTEXT_TEST") != "" {
		return "test"
	}
	return "default"
}

// UserDirectoryPath returns the MWCLIDIR in the user home directory (or similar directory) that can be used for storage
func UserDirectoryPath() string {
	return dirs.UserDirectoryPath(MWCLIDIR)
}

// UserDirectoryPathForCmd is a path within the application directory for the user that can be used for storage for the command
func UserDirectoryPathForCmd(cmdName string) string {
	return dirs.UserDirectoryPath(MWCLIDIR + string(os.PathSeparator) + cmdName)
}
