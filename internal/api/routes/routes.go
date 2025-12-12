package routes

import (
    "api.teklifYonetimi/internal/api/handlers"

    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    // Test endpoint
    r.GET("/ping", handlers.PingHandler)

    // Company routes
    RegisterCompanyRoutes(r)
}
