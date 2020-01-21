package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Golang Spell version number",
	Long:  `Shows the Golang Spell current installed version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Golang Spell v1.0.0 -- HEAD")
	},
}
