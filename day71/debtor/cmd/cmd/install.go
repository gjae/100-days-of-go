package cmd

import (
	"os"
	"time"

	"github.com/gjae/debtor/internal/repository"
	"github.com/gjae/debtor/internal/utils"
	"github.com/spf13/cobra"
)

var force bool

var InstallCommand = &cobra.Command{
	Use:   "install",
	Short: "Verified if exists .db file, in case that not exists then will be created",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.OpenFile(".debtor.db", os.O_APPEND, os.ModeAppend); err == nil {
			if force {
				os.Remove(".debtor.db")
				<-time.After(time.Second)
			} else {
				return
			}
		}
		db := utils.DB()
		repository.RunMigration(db)

	},
}
