package cmd

import (
	"github.com/danilovalente/golangspell/domain"
	"github.com/danilovalente/golangspell/usecase"
	"github.com/spf13/cobra"
)

func buildUpdateCommand() *domain.Command {
	return &domain.Command{
		Name:             "update",
		ShortDescription: "The update command updates the golangspell tool",
		LongDescription: `The update command updates the golangspell tool

Syntax: 
golangspell update
`,
	}
}

func runUpdateCommand(cmd *cobra.Command, args []string) {
	usecase.Update()
}
