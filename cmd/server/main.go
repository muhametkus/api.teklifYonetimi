package main

import (
    "api.teklifYonetimi/internal/api/routes"
    "api.teklifYonetimi/internal/config"
    "api.teklifYonetimi/internal/database"
    _ "api.teklifYonetimi/docs"

    "github.com/gin-gonic/gin"
)

// @title           Teklif Yönetimi API
// @version         1.0
// @description     Teklif Yönetim Sistemi API Dokümantasyonu
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
    cfg := config.LoadConfig()

    database.Connect()
	
	database.RunMigrations()

    r := gin.Default()

    routes.RegisterRoutes(r)

    r.Run(":" + cfg.ServerPort)
}
