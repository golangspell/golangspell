package domain

import (
	"github.com/danilovalente/golangspell/appcontext"
	"github.com/danilovalente/golangspell/cmd"
	"github.com/danilovalente/golangspell/config"
	"strings"
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
	Name             string          `json:"name"`
	ShortDescription string          `json:"shortDescription"`
	LongDescription  string          `json:"longDescription"`
	Flags            map[string]Flag `json:"flags"`
	ValidArgs        []string        `json:"validArgs"`
}

//Spell maps a Golangspell plugin
type Spell struct {
	Name      string             `json:"name"`
	URL       string             `json:"url"`
	Commands  map[string]Command `json:"commands"`
	Installed bool               `json:"installed"`
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
		Author:  cmd.Author,
		License: cmd.UserLicense,
		DefaultSpells: []GolangLibrary{
			{
				URL: "https://github.com/danilovalente/golangspell-core", Name: "golangspell-core",
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
