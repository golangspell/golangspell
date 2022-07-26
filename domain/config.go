package domain

import (
	"strings"

	"github.com/golangspell/golangspell/appcontext"
	"github.com/golangspell/golangspell/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

//Flag s defines special behaviors and configurations to the commands
type Flag struct {
	Name      string `json:"name"`
	Shorthand string `json:"shorthand"`
	Value     string `json:"value"`
	Usage     string `json:"usage"`
}

//Command is an available command in a specific Spell (plugin)
type Command struct {
	Name             string           `json:"name"`
	ShortDescription string           `json:"shortDescription"`
	LongDescription  string           `json:"longDescription"`
	Flags            map[string]*Flag `json:"flags"`
	ValidArgs        []string         `json:"validArgs"`
}

//RunCommandFunction specifies a function for running a command
type RunCommandFunction func(cmd *cobra.Command, args []string)

//CobraCommand creates a cobra.Command from the domain.Command specification
func (command *Command) CobraCommand(runCommandFunction RunCommandFunction) *cobra.Command {
	spellCMD := &cobra.Command{
		Use:       command.Name,
		Short:     command.ShortDescription,
		Long:      command.LongDescription,
		Run:       runCommandFunction,
		ValidArgs: command.ValidArgs,
	}
	for _, flag := range command.Flags {
		spellCMD.PersistentFlags().AddFlag(&pflag.Flag{Name: flag.Name, Shorthand: flag.Shorthand, Usage: flag.Usage})
	}
	return spellCMD
}

//Spell maps a Golangspell plugin
type Spell struct {
	Name      string              `json:"name"`
	URL       string              `json:"url"`
	Commands  map[string]*Command `json:"commands"`
	Installed bool                `json:"installed"`
	Version   string              `json:"version"`
}

//URLToPackage returns the package name referenced by the URL
func (spell *Spell) URLToPackage() string {
	return strings.ReplaceAll(strings.ReplaceAll(spell.URL, "http://", ""), "https://", "")
}

//Config holds the Golangspell tool configuration
type Config struct {
	Author        string            `json:"author"`
	License       string            `json:"license"`
	DefaultSpells []GolangLibrary   `json:"defaultSpells"`
	Spellbook     map[string]*Spell `json:"spellbook"`
}

/*
ConfigRepository defines the repository capabilities that should be found in a Repository implementation for Config
*/
type ConfigRepository interface {
	appcontext.Component
	Get() (*Config, error)
	Save(config *Config) (string, error)
}

//BuildDefaultConfig used to bootstrap application at first execution
func BuildDefaultConfig() Config {
	return Config{
		Author:  config.Author,
		License: config.UserLicense,
		DefaultSpells: []GolangLibrary{
			{
				URL: "https://github.com/golangspell/golangspell-core", Name: "golangspell-core",
			},
		},
	}
}

//InitConfig lazily loads a Config
func InitConfig() appcontext.Component {
	configRepository := appcontext.Current.Get(appcontext.ConfigRepository).(ConfigRepository)
	config, err := configRepository.Get()
	if err != nil {
		panic(err)
	}
	return config
}

//GetConfig from the Current Application Context
func GetConfig() *Config {
	return appcontext.Current.Get(appcontext.Config).(*Config)
}

//GetConfigRepository from the Current Application Context
func GetConfigRepository() ConfigRepository {
	return appcontext.Current.Get(appcontext.ConfigRepository).(ConfigRepository)
}

func init() {
	if config.Values.TestRun {
		return
	}

	appcontext.Current.Add(appcontext.Config, InitConfig)
}
