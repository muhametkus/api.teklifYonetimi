package handlers

import (
	"api.teklifYonetimi/internal/database"
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	sqlDB, err := database.DB.DB()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "DB object error",
			"error":   err.Error(),
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(500, gin.H{
			"message": "DB ping failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "pong",
		"db":      "connected",
	})
}
