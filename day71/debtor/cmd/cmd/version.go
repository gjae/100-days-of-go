package cmd

import (
	"fmt"

	"github.com/gjae/debtor/internal/utils"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print current app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(utils.VERSION)
	},
}
