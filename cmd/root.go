package cmd

import (
	"fmt"

	"github.com/danilovalente/golangspell/config"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	CfgFile     string
	UserLicense string
	Author      string
	//DefautConfigFile holds the Golangspell's config file path
	DefautConfigFile string = config.ConfigFilePath

	RootCmd = &cobra.Command{
		Use:   "golangspell",
		Short: "Golang Spell - A Golang Code generator for building Microservices",
		Long: `Golang Spell is a CLI library for Go. 
Golang Spell makes it possible to build lightning fast Microservices in Go 
in an easy and productive way.
Welcome to the tool that will kick out the boilerplate code 
and drive you through new amazing possibilities`,
	}
)

// Execute executes the root command.
func Execute() error {
	return RootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", fmt.Sprintf("config file (default is %s)", DefautConfigFile))
	RootCmd.PersistentFlags().StringVarP(&Author, "author", "a", "", "author name for copyright attribution")
	RootCmd.PersistentFlags().StringVarP(&UserLicense, "license", "l", "Apache", "name of license for the project")
	RootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("config", DefautConfigFile)
	viper.SetDefault("license", "Apache")
}

func initConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
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
