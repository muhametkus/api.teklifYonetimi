package database

import (
    "fmt"
    "log"
    "api.teklifYonetimi/internal/config"
    "api.teklifYonetimi/internal/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    host := config.GetEnv("DB_HOST")
    port := config.GetEnv("DB_PORT")
    user := config.GetEnv("DB_USER")
    pass := config.GetEnv("DB_PASS")
    name := config.GetEnv("DB_NAME")
    ssl  := config.GetEnv("DB_SSL")

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        host, user, pass, name, port, ssl,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Database bağlantı hatası: %v", err)
    }

    log.Println("✅ Database bağlantısı başarılı!")
    DB = db

    migrate()
}

func migrate() {
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
