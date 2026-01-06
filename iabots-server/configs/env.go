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
	AppEnv  string
	AppPort string
	DB      DBConfig
}

var Env EnvConfig

func LoadEnv() {
	Env = EnvConfig{
		AppEnv:  getEnv("APP_ENV", "local"),
		AppPort: firstEnv("APP_PORT", "PORT", "5001"),
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "iabots"),
			Password: getEnv("DB_PASSWORD", "123456"),
			Name:     getEnv("DB_NAME", "iabots_local"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}
	log.Println("✅ Environment variables loaded successfully")
}

func firstEnv(primary, secondary, fallback string) string {
	if v := os.Getenv(primary); v != "" {
		return v
	}
	if v := os.Getenv(secondary); v != "" {
		return v
	}
	return fallback
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
