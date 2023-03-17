package main

import (
	"github.com/gjae/debtor/cmd/cmd"
	"github.com/gjae/debtor/internal/repository"
	"gorm.io/gorm"
)

func startMigrations(db *gorm.DB) {
	repository.RunMigration(db)
}

func main() {
	cmd.Executable()
}
