package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define a route that responds with "Hello, World!" when accessed
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Run the server on port 8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
