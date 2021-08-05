package cmd

import (
	"fmt"

	"github.com/golangspell/golangspell/appcontext"
	"github.com/golangspell/golangspell/config"
	"github.com/spf13/cobra"
)

//GetRootCmd lazily loads a RootCmd to start the CLI application
func GetRootCmd() appcontext.Component {
	return &cobra.Command{
		Use:   "golangspell",
		Short: "Golang Spell - A Golang Code generator for building Microservices",
		Long: `Golang Spell is a CLI library for Go. 
Golang Spell makes it possible to build lightning fast Microservices in Go 
in an easy and productive way.
Welcome to the tool that will kick out the boilerplate code 
and drive you through new amazing possibilities`,
	}
}

// Execute executes the root command.
func Execute() error {
	rootCmd := appcontext.Current.Get(appcontext.RootCmd).(*cobra.Command)
	return rootCmd.Execute()
}

func addInternalCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(
		buildAddspellCommand().CobraCommand(runAddspellCommand))
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(
		buildUpdatespellCommand().CobraCommand(runUpdatespellCommand))
	rootCmd.AddCommand(
		buildRemovespellCommand().CobraCommand(runRemovespellCommand))
	rootCmd.AddCommand(
		buildUpdateCommand().CobraCommand(runUpdateCommand))
}

func init() {
	cobra.OnInitialize(initConfig)

	appcontext.Current.Add(appcontext.RootCmd, GetRootCmd)
	rootCmd := appcontext.Current.Get(appcontext.RootCmd).(*cobra.Command)
	rootCmd.PersistentFlags().StringVar(&config.CfgFile, "config", "", fmt.Sprintf("config file (default is %s)", config.DefautConfigFile))
	rootCmd.PersistentFlags().StringVarP(&config.Author, "author", "a", "", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&config.UserLicense, "license", "l", "Apache", "name of license for the project")
	addInternalCommands(rootCmd)
}

func initConfig() {
}
