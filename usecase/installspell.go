package usecase

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danilovalente/golangspell/domain"
)

//InstallSpell using go get
func InstallSpell(golangLibrary *domain.GolangLibrary, config *domain.Config) error {
	fmt.Printf("Installing Spell %s ...\n", golangLibrary.Name)
	execCmd := exec.Command("go", "get", golangLibrary.URLToPackage())
	execCmd.Env = os.Environ()
	execCmd.Env = append(execCmd.Env, "GO111MODULE=off")
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	err := execCmd.Run()
	if err != nil {
		return fmt.Errorf("go get %s failed with %s", golangLibrary.URLToPackage(), err.Error())
	}
	fmt.Printf("Spell %s installed!\n", golangLibrary.Name)
	return nil
}
