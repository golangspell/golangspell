package main

import (
	"fmt"

	"github.com/danilovalente/golangspell/usecase"

	"github.com/danilovalente/golangspell/cmd"
	_ "github.com/danilovalente/golangspell/gateway/filesystem"
)

func main() {
	usecase.LoadConfig()
	usecase.LoadSpells()
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("An error occurred while executing the command. Message: %s\n", err.Error())
	}
}
