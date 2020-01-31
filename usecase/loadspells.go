package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danilovalente/golangspell/appcontext"
	"github.com/danilovalente/golangspell/domain"
	"github.com/spf13/cobra"
)

func loadSpellCommand(spell *domain.Spell, command *domain.Command) {
	spellCMD := command.CobraCommand(
		func(cmd *cobra.Command, args []string) {
			library := domain.GolangLibrary{Name: spell.Name, URL: spell.URL}
			spellCommand := exec.Command(library.BinPath(), append([]string{command.Name}, args...)...)
			spellCommand.Stdout = os.Stdout
			spellCommand.Stderr = os.Stderr
			err := spellCommand.Run()
			if err != nil {
				log.Fatalf("%s failed with %s\n", command.Name, err)
			}
		})

	rootCmd := appcontext.Current.Get(appcontext.RootCmd).(*cobra.Command)
	rootCmd.AddCommand(spellCMD)
}

func loadSpellDescription(golangLibrary *domain.GolangLibrary, config *domain.Config) {
	fmt.Printf("Loading Spell %s description ...\n", golangLibrary.Name)
	execCmd := exec.Command(golangLibrary.BinPath(), "build-config")
	out, err := execCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("%s build-config failed with %s\n", golangLibrary.BinPath(), err)
	}
	var spell domain.Spell
	err = json.Unmarshal(out, &spell)
	if err != nil {
		panic(err)
	}
	if nil == config.Spellbook {
		config.Spellbook = make(map[string]*domain.Spell, 0)
	}
	spell.Installed = true
	config.Spellbook[spell.Name] = &spell
	repo := domain.GetConfigRepository()
	repo.Save(config)
	fmt.Printf("Spell %s description loaded\n", golangLibrary.Name)
}

//LoadSpells and configure cmds to call them
func LoadSpells() {
	fmt.Println("Loading Spells ...")
	config := domain.GetConfig()
	for _, golangLibrary := range config.DefaultSpells {
		if nil == config.Spellbook || nil == config.Spellbook[golangLibrary.Name] {
			AddSpell(&golangLibrary, config)
		}
	}
	for _, spell := range config.Spellbook {
		for _, command := range spell.Commands {
			if command.Name != "build-config" {
				loadSpellCommand(spell, command)
			}
		}
	}
	fmt.Println("Spells loaded!")
}
