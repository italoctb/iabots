package configs

import (
	"log"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type EnvConfig struct {
	DB      DBConfig
	AppPort string
}

var Env EnvConfig

func LoadEnv() {
	Env = EnvConfig{
		AppPort: getEnv("APP_PORT", "8080"),
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "admin"),
			Password: getEnv("DB_PASSWORD", "123456"),
			Name:     getEnv("DB_NAME", "whatsapp"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}
	log.Println("✅ Environment variables loaded successfully")
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
