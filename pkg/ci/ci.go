package ci

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// CreateJenkinsfile creates the jenkinsfile
func CreateJenkinsfile(path string) error {
	j := []byte(jenkinsfileTemplate)
	err := ioutil.WriteFile(path+"/Jenkinsfile", j, 0644)
	if err != nil {
		return errors.Wrap(err, "unable to write jenkinsfile")
	}

	return nil
}
