package main

import (
    "api.teklifYonetimi/internal/api/routes"
    "api.teklifYonetimi/internal/config"
    "api.teklifYonetimi/internal/database"

    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()

    database.Connect()

    r := gin.Default()

    routes.RegisterRoutes(r)

    r.Run(":8082")
}
