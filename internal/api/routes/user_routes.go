package routes

import (
	"api.teklifYonetimi/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userHandler := handlers.NewUserHandler()

	users := r.Group("/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("", userHandler.GetUsers)
	}
}
