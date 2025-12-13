package routes

import (
	"api.teklifYonetimi/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	authHandler := handlers.NewAuthHandler()

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}
}
