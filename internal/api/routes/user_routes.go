package routes

import (
	"api.teklifYonetimi/internal/api/handlers"
	"api.teklifYonetimi/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userHandler := handlers.NewUserHandler()

	users := r.Group("/users")
	users.Use(middleware.JWTAuthMiddleware())
	{
		// List & Detail (ADMIN + USER?) -> Belki sadece ADMIN?
		// Şimdilik herkese açık (kendi şirketindekileri görmeli aslında, service katmanında filtre lazım)
		users.GET("", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUserByID)

		// Create/Update/Delete -> ADMIN
		users.POST("", middleware.RequireRole("ADMIN"), userHandler.CreateUser)
		users.PUT("/:id", middleware.RequireRole("ADMIN"), userHandler.UpdateUser)
		users.DELETE("/:id", middleware.RequireRole("ADMIN"), userHandler.DeleteUser)
	}
}
