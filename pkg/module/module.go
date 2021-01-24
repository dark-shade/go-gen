package module

import (
	"os/exec"

	"github.com/pkg/errors"
)

// InitiateGoMod initializes go mod i.e. go mod init
func InitiateGoMod() error {
	cmd := exec.Command("go", "mod", "init")
	err := cmd.Run()
	if err != nil {
		return errors.Wrap(err, "unable to initialize go mod")
	}

	return nil
}
