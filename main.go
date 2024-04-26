package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "freed",
			"version":     "0.0.1",
			"status":      "up",
		})
	})

	r.Run(":42069")
}
