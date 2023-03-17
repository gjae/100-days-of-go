package models

import (
	"errors"

	"github.com/gjae/debtor/internal/utils"
	"gorm.io/gorm"
)

const (
	AccountEntry = iota
	AccountLoan
)

const (
	AccountStatusOpened = iota
	AccountStatusClose
)

type Client struct {
	gorm.Model
	Name     string `gorm:"default:''"`
	Lastname string `gorm:"default:''"`
	Ced      string `gorm:"default:''"`
	Accounts []Account
}

type Account struct {
	gorm.Model
	Type        int8
	ClientID    int
	Client      Client
	Amount      float32 `gorm:"type:decimal(9,2)"`
	Description string  `gorm:"size:20"`
	Payments    []Payment
	Status      uint8
}

type Payment struct {
	gorm.Model
	Account   Account
	AccountID int
	Amount    float32 `gorm:"type:decimal(9,2)"`
}

func NewClient(name, lastname, ced string) *Client {
	return &Client{
		Name:     name,
		Lastname: lastname,
		Ced:      ced,
	}
}

var (
	ErrClientNotFound = errors.New("User not was found")
)

func (c *Client) Exists() bool {
	db := utils.DB()
	var client *Client
	db.First(&client, "ced = ?", c.Ced)

	return client.ID != 0
}

func FindClient(ced string) (*Client, error) {
	var client *Client
	db := utils.DB()

	db.First(&client, "ced = ?", ced)

	if client == nil {
		return nil, ErrClientNotFound
	}

	return client, nil

}

func (c *Client) Insert() uint {
	db := utils.DB()
	db.Create(c)

	return c.ID
}

func (c *Client) NewAccount(accountType int8, amount float32, description string) *Account {
	db := utils.DB()
	account := &Account{
		Type:        accountType,
		Amount:      amount,
		Description: description,
		Status:      AccountStatusOpened,
		ClientID:    int(c.ID),
		Client:      *c,
	}

	db.Model(&Account{}).Create(account)

	return account
}

// Accounts functions

func ClientHasOpenedAccount(client *Client) bool {
	db := utils.DB()
	var account *Account

	db.Model(&Account{}).Where("status = ?", AccountStatusOpened).Where("client_id = ?", client.ID).First(&account)

	return account.ID != 0
}
