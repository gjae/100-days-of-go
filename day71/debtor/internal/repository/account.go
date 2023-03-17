package repository

import (
	"errors"

	"github.com/fatih/color"
	"github.com/gjae/debtor/internal/utils"
	models "github.com/gjae/debtor/pkg"
)

var (
	ErrClientNotFound = errors.New("Client not was found")
)

func NewAccount(ced string, amount float32, description string, typeAccount int) (*models.Account, error) {
	c, err := models.FindClient(ced)

	if err != nil {
		color.Red("Client not found, first register them")
		return nil, ErrClientNotFound
	}

	if models.ClientHasOpenedAccount(c) {
		color.Red("This client has opened account")
		return nil, ErrClientNotFound
	}

	a := c.NewAccount(int8(typeAccount), amount, description)

	return a, nil
}

func ListAccounts(typeAccount int) []models.Account {
	db := utils.DB()
	var accounts []models.Account
	_ = db.Model(&models.Account{}).Where("status = ?", typeAccount).Joins("Client").Find(&accounts)

	return accounts
}
