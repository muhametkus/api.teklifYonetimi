package main

import (
	"api.teklifYonetimi/internal/api/routes"
	"api.teklifYonetimi/internal/config"
	"api.teklifYonetimi/internal/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Config yükleme
	cfg := config.LoadConfig()

	// Database bağlantısı
	db := database.InitDB(cfg)
	defer database.CloseDB(db)

	// Gin router oluşturma
	r := gin.Default()

	// Routes kaydetme
	routes.RegisterRoutes(r, db)

	// Server başlatma
	log.Printf("Server starting on port %s...", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
