package database

import (
	"api.teklifYonetimi/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func InitDB(cfg *config.Config) *sql.DB {
	// Şimdilik database bağlantısı olmadan çalışacak
	// İleride PostgreSQL bağlantısı eklenebilir
	log.Println("Database initialization skipped (not configured yet)")
	return nil
}

func CloseDB(db *sql.DB) {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}
}

// PostgreSQL bağlantı örneği (şimdilik kullanılmıyor)
func connectPostgreSQL(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connected successfully")
	return db, nil
}
