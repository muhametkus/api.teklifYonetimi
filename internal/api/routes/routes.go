package routes

import (
    "api.teklifYonetimi/internal/api/handlers"

    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {
    // Test endpoint
    r.GET("/ping", handlers.PingHandler)

    // Company routes
    RegisterCompanyRoutes(r)

    // User routes
    RegisterUserRoutes(r)

    // Auth routes
    RegisterAuthRoutes(r)

    // Quotation routes
    RegisterQuotationRoutes(r)

    // Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
