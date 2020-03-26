package usecase

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danilovalente/golangspell/domain"
)

//ShowSpellVersion shows the version of an specified Spell
func ShowSpellVersion(golangLibrary *domain.GolangLibrary, config *domain.Config) error {
	execCmd := exec.Command(golangLibrary.BinPath(), golangLibrary.Name+"-version")
	execCmd.Env = os.Environ()
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	err := execCmd.Run()
	if err != nil {
		return fmt.Errorf("golangspell version %s failed. Error: %s", golangLibrary.Name, err.Error())
	}
	return nil
}
