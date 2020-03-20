package usecase

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danilovalente/golangspell/domain"
)

//Update GolangSpell tool
func Update() error {
	fmt.Printf("updating GolangSpell ...\n")
	library := domain.GolangLibrary{Name: "golangspell", URL: "https://github.com/danilovalente/golangspell"}
	err := os.RemoveAll(library.SrcPath())
	if err != nil {
		return fmt.Errorf("An error occurred while trying to remove the spell: %s", err.Error())
	}

	err = os.Remove(library.BinPath())
	if err != nil {
		return fmt.Errorf("An error occurred while trying to remove the spell: %s", err.Error())
	}

	execCmd := exec.Command("go", "get", "github.com/danilovalente/golangspell")
	execCmd.Env = os.Environ()
	execCmd.Env = append(execCmd.Env, "GO111MODULE=off")
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	err = execCmd.Run()
	if err != nil {
		return fmt.Errorf("go get %s failed with %s", "github.com/danilovalente/golangspell", err.Error())
	}
	fmt.Printf("GolangSpell updated!\n")
	return nil
}
