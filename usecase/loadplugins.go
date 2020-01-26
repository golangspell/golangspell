package usecase

import (
	"log"
	"os/exec"

	"github.com/danilovalente/golangspell/cmd"

	"github.com/danilovalente/golangspell/appcontext"
	"github.com/danilovalente/golangspell/domain"
	"github.com/spf13/cobra"
)

func init() {

}

func loadPluginCommand(spell domain.Spell, command domain.Command) {
	spellCMD := &cobra.Command{
		Use:   command.Name,
		Short: command.ShortDescription,
		Long:  command.LongDescription,
		Run: func(cmd *cobra.Command, args []string) {
			spellCommand := exec.Command(command.Name, args...)
			err := spellCommand.Run()
			if err != nil {
				log.Fatalf("%s failed with %s\n", command.Name, err)
			}
		},
	}
	cmd.RootCmd.AddCommand(spellCMD)
}

//LoadPlugins and configure cmds to call them
func LoadPlugins() {
	config := appcontext.Current.Get(appcontext.Config).(domain.Config)
	for _, spell := range config.Spellbook {
		if !spell.Installed {
			InstallPlugin(spell)
		}
		for _, command := range spell.Commands {
			loadPluginCommand(spell, command)
		}
	}
}
