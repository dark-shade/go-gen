package dir

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// CreateDirectoryStructure creates the directory structure for the project
func CreateDirectoryStructure(path string) []error {
	var finalError []error

	err := os.Mkdir(path+"/bin", 0755)
	if err != nil {
		finalError = append(finalError, errors.Wrap(err, "unable to create the bin directory"))
	}

	err = os.Mkdir(path+"/docs", 0755)
	if err != nil {
		finalError = append(finalError, errors.Wrap(err, "unable to create the docs directory"))
	}

	err = os.Mkdir(path+"/pkg", 0755)
	if err != nil {
		finalError = append(finalError, errors.Wrap(err, "unable to create the pkg directory"))
	}

	ver := []byte("0.1.0")
	err = ioutil.WriteFile(path+"/VERSION", ver, 0644)
	if err != nil {
		finalError = append(finalError, errors.Wrap(err, "unable to create VERSION file"))
	}

	m := []byte(`
	package main

	func main() {

	}
	`)
	err = ioutil.WriteFile(path+"/main.go", m, 0644)
	if err != nil {
		finalError = append(finalError, errors.Wrap(err, "unable to create VERSION file"))
	}

	return nil
}
