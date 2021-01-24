package dockerfile

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// CreateDockerfile creates the dockerfile
func CreateDockerfile(path string) error {
	m := []byte(dockerfileTemplate)
	err := ioutil.WriteFile(path+"/Dockerfile", m, 0644)
	if err != nil {
		return errors.Wrap(err, "unable to write dockerfile")
	}

	return nil
}
