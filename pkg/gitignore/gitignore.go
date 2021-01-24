package gitignore

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// CreateGitIgnore creates the gitignore file
func CreateGitIgnore(path string) error {
	g := []byte(gitignoreTemplate)
	err := ioutil.WriteFile(path+"/.gitignore", g, 0644)
	if err != nil {
		return errors.Wrap(err, "unable to write gitignore")
	}

	return nil
}
