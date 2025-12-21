package database

import (
	"fmt"
	"log"

	"api.teklifYonetimi/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect database bağlantısını kurar
func Connect() {
	host := config.GetEnv("DB_HOST")
	port := config.GetEnv("DB_PORT")
	user := config.GetEnv("DB_USER")
	pass := config.GetEnv("DB_PASS")
	name := config.GetEnv("DB_NAME")
	ssl := config.GetEnv("DB_SSL")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, pass, name, port, ssl,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Database bağlantı hatası: %v", err)
	}

	DB = db
	log.Println("✅ Database bağlantısı başarılı!")
}
