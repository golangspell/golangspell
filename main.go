package main

import (
	"fmt"

	"github.com/golangspell/golangspell/usecase"

	"github.com/golangspell/golangspell/cmd"
	_ "github.com/golangspell/golangspell/gateway/filesystem"
)

func main() {
	usecase.LoadConfig()
	usecase.LoadSpells()
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("An error occurred while executing the command. Message: %s\n", err.Error())
	}
}
