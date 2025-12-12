package routes

import (
    "github.com/gin-gonic/gin"
    "api.teklifYonetimi/internal/api/handlers"
)

func RegisterRoutes(r *gin.Engine) {
    // Basit test endpoint
    r.GET("/ping", handlers.PingHandler)
}
