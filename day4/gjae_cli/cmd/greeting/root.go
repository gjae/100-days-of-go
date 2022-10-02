package greeting

import (
	"fmt"
	"os"
   
	"github.com/spf13/cobra"
)

var version string = "0.0.1"

var rootCommand = &cobra.Command {
	Use: "greeting",
	Short: "Greeting - a greeting beautiful app",
	Long: "Greeting - a greeting beautiful app cli with Golang",
	Version: version,
	Run: func(command *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println("Command not found")
		os.Exit(1)
	}
}