package cmd

import (
	"fmt"

	"github.com/danilovalente/golangspell/appcontext"
	"github.com/danilovalente/golangspell/config"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func init() {
	cobra.OnInitialize(initConfig)

	appcontext.Current.Add(appcontext.RootCmd, GetRootCmd)
	rootCmd := appcontext.Current.Get(appcontext.RootCmd).(*cobra.Command)
	rootCmd.PersistentFlags().StringVar(&config.CfgFile, "config", "", fmt.Sprintf("config file (default is %s)", config.DefautConfigFile))
	rootCmd.PersistentFlags().StringVarP(&config.Author, "author", "a", "", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&config.UserLicense, "license", "l", "Apache", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.AddCommand(
		buildAddspellCommand().CobraCommand(runAddspellCommand))
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(
		buildUpdatespellCommand().CobraCommand(runUpdatespellCommand))
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("config", config.DefautConfigFile)
	viper.SetDefault("license", "Apache")
}

func initConfig() {
	if config.CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(config.CfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".golangspell")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
