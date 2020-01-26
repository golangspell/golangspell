package domain

import (
	"github.com/danilovalente/golangspell/appcontext"
	"github.com/danilovalente/golangspell/cmd"
	"github.com/danilovalente/golangspell/config"
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

//Config holds the Golangspell tool configuration
type Config struct {
	Author    string           `json:"author"`
	License   string           `json:"license"`
	Spellbook map[string]Spell `json:"spellbook"`
}

/*
ConfigRepository defines the repository capabilities that should be found in a Repository implementation for Config
*/
type ConfigRepository interface {
	appcontext.Component
	Get() (*Config, error)
	Save(config *Config) (string, error)
}

func BuildDefaultConfig() Config {
	core := Spell{
		Name: "core",
		URL:  "https://github.com/danilovalente/golangspell-core",
		Commands: map[string]Command{
			"init": Command{
				Name:             "init",
				ShortDescription: "The init command creates a new Golang application using the Golangspell base structure",
				LongDescription: `The init command creates a new Golang application using the Golangspell base structure
The Architectural Model is based in the Clean Architecture and is the basis to add more resources like domain models and repositories.
Args:
module: Module name (required) to initialize with 'Go Modules'. Example: mydomain.com/myapplication"
appname: App name (required) to initialize with 'Go Modules'. Example: myapplication

Syntax: 
golangspell init [module] [appname]
`,
				ValidArgs: []string{"module", "name"},
			},
		},
	}
	return Config{
		Author:    cmd.Author,
		License:   cmd.UserLicense,
		Spellbook: map[string]Spell{"core": core},
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

func init() {
	if config.Values.TestRun {
		return
	}

	appcontext.Current.Add(appcontext.Config, InitConfig)
}
