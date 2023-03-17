package repository

import (
	"log"

	"github.com/gjae/debtor/internal/utils"
	models "github.com/gjae/debtor/pkg"
)

func NewClient(name, lastname, ced string) (*models.Client, error) {
	client := models.NewClient(name, lastname, ced)
	exists := client.Exists()

	if exists {
		log.Printf("The client with ced %s already exists\n", ced)
		return nil, ErrClientExists
	}

	_ = client.Insert()

	return client, nil
}

func ClientExists(ced string) (*models.Client, error) {
	client, err := models.FindClient(ced)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func ListClients() []*models.Client {
	db := utils.DB()

	var clients []*models.Client
	db.Find(&clients)

	return clients
}
