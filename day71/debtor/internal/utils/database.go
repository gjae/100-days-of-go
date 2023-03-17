package utils

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var once sync.Once

type single struct {
	db *gorm.DB
}

var instance *single

func getInstance() *single {
	if instance == nil {
		once.Do(
			func() {
				db, err := gorm.Open(sqlite.Open(".debtor.db"), &gorm.Config{
					Logger: nil,
				})

				if err != nil {
					log.Fatalf("Opening database has been failed: %v", err.Error())
				}

				instance = &single{db: db}
			},
		)
	} else {
		log.Println("An instance of gorm DATABASE was created")
	}

	return instance
}

func DB() *gorm.DB {
	return getInstance().db
}
