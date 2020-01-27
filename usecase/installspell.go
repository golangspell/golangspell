package usecase

import (
	"fmt"
	"os/exec"

	"github.com/danilovalente/golangspell/domain"
)

//InstallSpell using go get
func InstallSpell(golangLibrary *domain.GolangLibrary, config *domain.Config) error {
	fmt.Printf("Installing Spell %s ...\n", golangLibrary.Name)
	execCmd := exec.Command("go", "get", golangLibrary.URLToPackage())
	outputBytes, err := execCmd.Output()
	if err != nil {
		return fmt.Errorf("go get %s failed with %s", golangLibrary.URLToPackage(), err.Error())
	}
	fmt.Println(string(outputBytes))
	fmt.Printf("Spell %s installed!\n", golangLibrary.Name)
	return nil
}
