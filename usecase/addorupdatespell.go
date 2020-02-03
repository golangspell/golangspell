package usecase

import (
	"fmt"

	"github.com/danilovalente/golangspell/domain"
)

//AddOrUpdateSpell to the Golangspell platform
func AddOrUpdateSpell(golangLibrary *domain.GolangLibrary, config *domain.Config) {
	err := InstallSpell(golangLibrary, config)
	if err != nil {
		fmt.Printf("An error occurred while trying to install the spell: %s\n", err.Error())
	} else {
		loadSpellDescription(golangLibrary, config)
	}

}
