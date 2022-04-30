package database

import (
	"app/server/database/migrations"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	uri := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		port := os.Getenv("DATABASE_PORT")
		url := os.Getenv("DATABASE_HOST")
		user := os.Getenv("DATABASE_USER")
		password := os.Getenv("DATABASE_PASSWORD")
		dbname := os.Getenv("DATABASE_NAME")
		str := "host=" + url + " port=" + port + " user=" + user + " dbname=" + dbname + " sslmode=disable password=" + password
		database, err = gorm.Open(postgres.Open(str), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
