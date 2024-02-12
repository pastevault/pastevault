package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the database connection
var DB *gorm.DB

// Connect connects to the database
func Connect() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("./pastes.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
