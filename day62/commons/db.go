package commons

import (
	"errors"

	"github.com/gjae/go-todo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type singletonDatabase struct {
	db *gorm.DB
}

var db *singletonDatabase

func GetDb(settings config.Settings) (*gorm.DB, error) {
	if db != nil {
		return db.db, nil
	}
	db = new(singletonDatabase)

	dbAux, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  settings.GetDatabaseConfig(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.db = dbAux

	return db.db, nil
}

func GetCurrentDB() (*gorm.DB, error) {
	if db == nil {
		return nil, errors.New("No database init")
	}

	return db.db, nil
}
