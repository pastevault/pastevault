package main

import (
	_ "pastevault.com/internal/db"
	"pastevault.com/internal/expired"
	"pastevault.com/internal/router"
)

func main() {
	// Delete expired pastes
	go expired.Delete()

	// Start the server
	router.Router()
}
