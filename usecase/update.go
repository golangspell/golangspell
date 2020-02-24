package usecase

import (
	"fmt"
	"os"
	"os/exec"
)

//Update GolangSpell tool
func Update() error {
	fmt.Printf("updating GolangSpell ...\n")
	execCmd := exec.Command("go", "get", "github.com/danilovalente/golangspell")
	execCmd.Env = os.Environ()
	execCmd.Env = append(execCmd.Env, "GO111MODULE=off")
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	err := execCmd.Run()
	if err != nil {
		return fmt.Errorf("go get %s failed with %s", "github.com/danilovalente/golangspell", err.Error())
	}
	fmt.Printf("GolangSpell updated!\n")
	return nil
}
