package cmd

import (
	"fmt"

	"github.com/danilovalente/golangspell/domain"
	"github.com/danilovalente/golangspell/usecase"
	"github.com/spf13/cobra"
)

func buildAddspellCommand() *domain.Command {
	return &domain.Command{
		Name:             "addspell",
		ShortDescription: "The addspell command adds a Spell (plugin) to the golangspell tool",
		LongDescription: `The addspell command adds a Spell (plugin) to the golangspell tool
Args:
url: URL to the Spell's code repository (required). Example: https://github.com/danilovalente/golangspell-core"
name: Spell name (required). Example: golangspell-core

Syntax: 
golangspell addspell [url] [name]
`,
		ValidArgs: []string{"url", "name"},
	}
}

func runAddspellCommand(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println(`The command addspell requires exactly two parameters: url and name
Syntax: 
golangspell addspell [url] [name]`)
		return
	}
	golangLibrary := domain.GolangLibrary{URL: args[0], Name: args[1]}
	config := domain.GetConfig()
	if nil == config.Spellbook || nil == config.Spellbook[golangLibrary.Name] {
		usecase.AddSpell(&golangLibrary, config)
		fmt.Printf("Spell %s added.\n", golangLibrary.Name)
	}
}
