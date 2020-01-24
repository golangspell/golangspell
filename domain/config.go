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

func BuildDefaultConfig() Config {
	core := Spell{}
	return Config{
		Author:    cmd.Author,
		License:   cmd.UserLicense,
		Spellbook: []Spell{core},
	}
}
