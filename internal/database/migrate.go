package database

import (
	"log"

	"api.teklifYonetimi/internal/models"
)

// RunMigrations database modellerini migrate eder
func RunMigrations() {
	if DB == nil {
		log.Fatal("❌ Migration için DB bağlantısı yok")
	}

	err := DB.AutoMigrate(
		&models.Company{},
		&models.User{},
		&models.Quotation{},
		&models.QuotationItem{},
	)
	if err != nil {
		log.Fatalf("❌ Migration hatası: %v", err)
	}

	log.Println("✅ Modeller başarıyla migrate edildi!")
}
