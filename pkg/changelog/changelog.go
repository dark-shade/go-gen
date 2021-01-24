package changelog

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// CreateChangeLog creates the changelog file
func CreateChangeLog(path string) error {
	c := []byte(changelogTemplate)
	err := ioutil.WriteFile(path+"/CHANGELOG.md", c, 0644)
	if err != nil {
		return errors.Wrap(err, "unable to write changelog")
	}

	return nil
}
