package database

import (
	"fmt"
	"log"

	"iabots-server/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		configs.Env.DB.Host,
		configs.Env.DB.Port,
		configs.Env.DB.User,
		configs.Env.DB.Password,
		configs.Env.DB.Name,
		configs.Env.DB.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("✅ Database connected successfully")
	return &Database{DB: db}, nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.DB
}
