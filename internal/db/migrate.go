package db

import "pastevault.com/internal/paste"

func migrate() {
	DB.AutoMigrate(&paste.Paste{})
}
