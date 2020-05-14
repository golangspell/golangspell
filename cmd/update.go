package cmd

import (
	"fmt"

	"github.com/golangspell/golangspell/domain"
	"github.com/golangspell/golangspell/usecase"
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
	err := usecase.Update()
	if err != nil {
		fmt.Printf("An error occurred while trying to update the golangspell tool. Message: %s\n", err.Error())
	}
}
