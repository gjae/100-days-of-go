package greeting

import (
	"github.com/spf13/cobra"
	"github.com/gjae/gjae_cli/pkg/greeting"
)

var getRandomMessage bool

var greetingCommand = &cobra.Command {
	Use: "greetme",
	Aliases: []string{"g"},
	Short: "Say hello to user",
	Args: cobra.ExactArgs(1),
	Run: func(command *cobra.Command, args []string) {
		var name greeting.Name = greeting.Name(args[0])
		name.SayHello(getRandomMessage)
	},
}

func init() {
	greetingCommand.Flags().BoolVarP(&getRandomMessage, "random-message", "r", false, "Using random message")
	rootCommand.AddCommand(greetingCommand)
}