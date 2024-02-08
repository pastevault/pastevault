package router

import (
	"net/http"
	"pastevault.com/internal/paste"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	// API Version 1
	v1 := r.Group("/v1")

	// Check if the server is up
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Get a paste
	v1.GET("/paste/:uuid", paste.GetPasteHandler)

	// Create a new paste
	v1.POST("/paste", paste.NewPasteHandler)

	// Run the server
	if err := r.Run(); err != nil {
		panic(err)
	}
}
