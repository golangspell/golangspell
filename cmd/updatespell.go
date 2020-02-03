package cmd

import (
	"fmt"

	"github.com/danilovalente/golangspell/domain"
	"github.com/danilovalente/golangspell/usecase"
	"github.com/spf13/cobra"
)

func buildUpdatespellCommand() *domain.Command {
	return &domain.Command{
		Name:             "updatespell",
		ShortDescription: "The updatespell command updates a Spell (plugin) in the golangspell tool",
		LongDescription: `The updatespell command updates a Spell (plugin) in the golangspell tool
Args:
url: URL to the Spell's code repository (required). Example: https://github.com/danilovalente/golangspell-core"
name: Spell name (required). Example: golangspell-core

Syntax: 
golangspell updatespell [url] [name]
`,
		ValidArgs: []string{"url", "name"},
	}
}

func runUpdatespellCommand(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println(`The command updatespell requires exactly two parameters: url and name
Syntax: 
golangspell updatespell [url] [name]`)
		return
	}
	golangLibrary := domain.GolangLibrary{URL: args[0], Name: args[1]}
	config := domain.GetConfig()
	usecase.AddOrUpdateSpell(&golangLibrary, config)
	fmt.Printf("Spell %s updated.\n", golangLibrary.Name)
}
