package main

import (
	"freed/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := database.New("./freed.db")

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	r := gin.Default()

	// Middlewares
	r.Use(gin.Recovery())

	// Routes
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "freed",
			"version":     "0.0.1",
			"status":      "up",
		})
	})

	r.Run(":42069")
}
