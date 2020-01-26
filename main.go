package main

import (
	"github.com/danilovalente/golangspell/usecase"

	"github.com/danilovalente/golangspell/cmd"
	_ "github.com/danilovalente/golangspell/gateway/filesystem"
)

func main() {
	usecase.LoadConfig()
	usecase.LoadPlugins()
	cmd.Execute()
}
