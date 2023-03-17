package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/fatih/color"

	models "github.com/gjae/debtor/pkg"
	"gorm.io/gorm"
)

var (
	ErrClientExists = errors.New("Client already exists")
)

func RunMigration(db *gorm.DB) {
	success := color.New(color.FgGreen, color.Bold)

	fmt.Print("Migrating Client model ...")
	db.AutoMigrate(&models.Client{})
	<-time.After(time.Second)
	success.Println("\rClient migration was been runned!")

	fmt.Print("Migrating Account model ...")
	db.AutoMigrate(&models.Account{})
	<-time.After(time.Second)
	success.Println("\rAccount migration was been runned!")

	fmt.Print("Migrating Payments model ...")
	db.AutoMigrate(&models.Payment{})
	<-time.After(time.Second)
	success.Println("\rPayments migration was been runned!")

}
