package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    ServerPort string
    DBHost     string
    DBPort     string
    DBUser     string
    DBPass     string
    DBName     string
    DBSSLMode  string
}

func LoadConfig() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system variables")
    }

    return &Config{
        ServerPort: getEnvOrDefault("SERVER_PORT", "8082"),
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBUser:     os.Getenv("DB_USER"),
        DBPass:     os.Getenv("DB_PASS"),
        DBName:     os.Getenv("DB_NAME"),
        DBSSLMode:  getEnvOrDefault("DB_SSL", "disable"),
    }
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system variables")
    }
}

func GetEnv(key string) string {
    return os.Getenv(key)
}
