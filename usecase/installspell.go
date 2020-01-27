package usecase

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/danilovalente/golangspell/domain"
)

//InstallSpell using go get
func InstallSpell(spell *domain.Spell, config *domain.Config) {
	fmt.Printf("Installing Spell %s ...\n", spell.Name)
	execCmd := exec.Command("go", "get", spell.URLToPackage())
	outputBytes, err := execCmd.Output()
	if err != nil {
		log.Fatalf("go get %s failed with %s\n", spell.URLToPackage(), err)
	}
	fmt.Println(string(outputBytes))
	spell.Installed = true
	repo := domain.GetConfigRepository()
	repo.Save(config)
	fmt.Printf("Spell %s installed!\n", spell.Name)
}
