package usecase

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danilovalente/golangspell/domain"
)

//RemoveSpell to the Golangspell platform
func RemoveSpell(spellName string, config *domain.Config) error {
	if nil == config.Spellbook || nil == config.Spellbook[spellName] {
		return fmt.Errorf("Spell %s not found", spellName)
	}
	spell := config.Spellbook[spellName]
	fmt.Printf("Removing Spell %s\n", spellName)
	library := domain.GolangLibrary{Name: spell.Name, URL: spell.URL}
	removeSpellCommand := exec.Command("rm", "-rf", library.SrcPath())
	removeSpellCommand.Stdout = os.Stdout
	removeSpellCommand.Stderr = os.Stderr
	err := removeSpellCommand.Run()
	if err != nil {
		return fmt.Errorf("An error occurred while trying to remove the spell: %s", err.Error())
	}

	removeSpellCommand = exec.Command("rm", library.BinPath())
	removeSpellCommand.Stdout = os.Stdout
	removeSpellCommand.Stderr = os.Stderr
	err = removeSpellCommand.Run()
	if err != nil {
		return fmt.Errorf("An error occurred while trying to remove the spell: %s", err.Error())
	}

	delete(config.Spellbook, spellName)
	repo := domain.GetConfigRepository()
	repo.Save(config)

	return nil
}
