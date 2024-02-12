package router

import (
	"net/http"
	"pastevault.com/internal/logger"
	"pastevault.com/internal/paste"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func Router() {
	r := gin.New()
	r.Use(sloggin.New(logger.Logger))
	r.Use(gin.Recovery())

	// API Version 1
	v1 := r.Group("/v1")

	// Check if the server is up
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Get raw paste
	v1.GET("/paste/:id/raw", paste.GetRawPasteHandler)

	// Run the server
	if err := r.Run(); err != nil {
		panic(err)
	}
}
