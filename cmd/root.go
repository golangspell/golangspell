package cmd

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	//DefautConfigFile holds the Golangspell's config file path
	DefautConfigFile string = "$HOME/.golangspell/.golangspell"
)

var (
	// Used for flags.
	CfgFile     string
	UserLicense string
	Author      string

	rootCmd = &cobra.Command{
		Use:   "gospell",
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
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", fmt.Sprintf("config file (default is %s)", DefautConfigFile))
	rootCmd.PersistentFlags().StringVarP(&Author, "author", "a", "", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&UserLicense, "license", "l", "Apache", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
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
