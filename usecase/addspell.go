package usecase

import (
	"fmt"

	"github.com/danilovalente/golangspell/domain"
)

//AddSpell to the Golangspell platform
func AddSpell(golangLibrary *domain.GolangLibrary, config *domain.Config) {
	err := InstallSpell(golangLibrary, config)
	if err != nil {
		fmt.Printf("An error occurred while trying to install the spell: %s\n", err.Error())
	} else {
		loadSpellDescription(golangLibrary, config)
	}

}
