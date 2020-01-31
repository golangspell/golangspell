package cmd

import (
	"fmt"

	"github.com/danilovalente/golangspell/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Golang Spell version number",
	Long:  `Shows the Golang Spell current installed version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Golang Spell v%s -- HEAD\n", config.Version)
	},
}
