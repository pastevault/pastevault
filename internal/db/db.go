package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the database connection
var DB *gorm.DB

// Connect connects to the database
func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./pastes.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migrate()
}
