package routes

import (
	"api.teklifYonetimi/internal/api/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	// Handlers
	pingHandler := handlers.NewPingHandler()

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		// Health check endpoints
		v1.GET("/ping", pingHandler.Ping)
		v1.GET("/hello", pingHandler.Hello)
	}

	// Root level endpoints
	r.GET("/hello", pingHandler.Hello)
}
