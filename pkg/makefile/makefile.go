package makefile

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// CreateMakefile creates the makefile
func CreateMakefile(path string) error {
	m := []byte(makefileTemplate)
	err := ioutil.WriteFile(path+"/Makefile", m, 0644)
	if err != nil {
		return errors.Wrap(err, "unable to write makefile")
	}

	return nil
}
