package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"} // Allow requests from Vue dev server
	r.Use(cors.New(config))

	// Define a route
	r.GET("/api/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Go backend!",
		})
	})

	// Start the server
	r.Run(":8000")
}
