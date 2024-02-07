package main

import (
	_ "pastevault.com/internal/db"
	"pastevault.com/internal/router"
)

func main() {
	router.Router()
}
