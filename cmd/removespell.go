package cmd

import (
	"fmt"

	"github.com/golangspell/golangspell/domain"
	"github.com/golangspell/golangspell/usecase"
	"github.com/spf13/cobra"
)

func buildRemovespellCommand() *domain.Command {
	return &domain.Command{
		Name:             "removespell",
		ShortDescription: "The removespell command removes a Spell (plugin) from the golangspell tool",
		LongDescription: `The removespell command removes a Spell (plugin) from the golangspell tool
Args:
name: Spell name (required). Example: golangspell-myspell

Syntax: 
golangspell removespell [name]
`,
		ValidArgs: []string{"name"},
	}
}

func runRemovespellCommand(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(`The command removespell requires exactly one parameter: name
Syntax: 
golangspell removespell [name]`)
		return
	}
	spellName := args[0]
	config := domain.GetConfig()
	err := usecase.RemoveSpell(spellName, config)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Spell %s removed.\n", spellName)
	}
}
