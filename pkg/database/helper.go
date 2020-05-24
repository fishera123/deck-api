package database

import (
	"github.com/jinzhu/gorm"
)

// todo: should get variables from .env
// todo: add migrations, auto migration with models???
func ConnectDb() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=database port=5432 user=deck_user dbname=deck_db password=verymuchsecret sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil

}
