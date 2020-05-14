package usecase

import (
	"fmt"
	"os"

	"github.com/golangspell/golangspell/domain"
)

//RemoveSpell to the Golangspell platform
func RemoveSpell(spellName string, config *domain.Config) error {
	if nil == config.Spellbook || nil == config.Spellbook[spellName] {
		return fmt.Errorf("Spell %s not found", spellName)
	}
	spell := config.Spellbook[spellName]
	fmt.Printf("Removing Spell %s\n", spellName)
	library := domain.GolangLibrary{Name: spell.Name, URL: spell.URL}
	err := os.RemoveAll(library.SrcPath())
	if err != nil {
		return fmt.Errorf("An error occurred while trying to remove the spell: %s", err.Error())
	}

	err = os.Remove(library.BinPath())
	if err != nil {
		return fmt.Errorf("An error occurred while trying to remove the spell: %s", err.Error())
	}

	delete(config.Spellbook, spellName)
	repo := domain.GetConfigRepository()
	_, err = repo.Save(config)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to update the spell configuration: %s", err.Error())
	}

	return nil
}
