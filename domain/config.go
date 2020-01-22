package domain

import "golangspell.com/golangspell/cmd"

//Flag s defines special behaviors and configurations to the commands
type Flag struct {
	Name      string `json:"name"`
	Shorthand string `json:"shorthand"`
	Value     string `json:"value"`
	Usage     string `json:"usage"`
}

//Command is an available command in a specific Spell (plugin)
type Command struct {
	Name  string `json:"name"`
	Flags []Flag `json:"flags"`
}

//Spell maps a Golangspell plugin
type Spell struct {
	Name     string    `json:"name"`
	URL      string    `json:"url"`
	Commands []Command `json:"commands"`
}

//Config holds the Golangspell tool configuration
type Config struct {
	Author    string  `json:"author"`
	License   string  `json:"license"`
	Spellbook []Spell `json:"spellbook"`
}

func buildCoreInitCommand() Command {
	coreInitFlagModule := Flag{
		Name:      "module",
		Shorthand: "m",
		Value:     "",
		Usage:     "Module name (required) to initialize with 'Go Modules'. Example: mydomain.com/myapplication",
	}
	coreInitFlagAppName := Flag{
		Name:      "appname",
		Shorthand: "n",
		Value:     "",
		Usage:     "App name (required) to initialize with 'Go Modules'. Example: myapplication",
	}
	coreInitFlags := make([]Flag, 2)
	coreInitFlags[0] = coreInitFlagModule
	coreInitFlags[1] = coreInitFlagAppName
	coreInit := Command{
		Name:  "init",
		Flags: coreInitFlags,
	}
	return coreInit
}

func BuildDefaultConfig() Config {
	coreCommands := make([]Command, 8)
	coreCommands[0] = buildCoreInitCommand()
	core := Spell{
		Name:     "core",
		URL:      "https://github.com/danilovalente/golangspell-core",
		Commands: coreCommands,
	}
	spellbook := make([]Spell, 1)
	spellbook[0] = core
	return Config{
		Author:    cmd.Author,
		License:   cmd.UserLicense,
		Spellbook: spellbook,
	}
}
