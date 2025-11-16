package main

import (
	"net/http"

	"github.com/aayushchugh/ayushchugh.com-api/config/env"
	"github.com/aayushchugh/ayushchugh.com-api/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	envConfig := env.LoadEnv()
	database.InitDB()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is working perfectly",
		})
	})

	r.Run(":" + envConfig.Port)
}
