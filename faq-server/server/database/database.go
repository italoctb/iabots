package database

import (
	"app/server/database/migrations"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var db *gorm.DB

var dbSql *sql.DB

func StartDB() {

	uri := os.Getenv("DATABASE_URL")
	var database *gorm.DB
	var databaseSql *sql.DB
	var err error
	var errSql error
	if uri != "" {
		fmt.Print(uri)
		database, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
		databaseSql, errSql = sql.Open("postgres", uri)
	} else {
		port := os.Getenv("DATABASE_PORT")
		url := os.Getenv("DATABASE_HOST")
		user := os.Getenv("DATABASE_USER")
		password := os.Getenv("DATABASE_PASSWORD")
		dbname := os.Getenv("DATABASE_NAME")
		str := "host=" + url + " port=" + port + " user=" + user + " dbname=" + dbname + " sslmode=disable password=" + password
		fmt.Print(str)
		database, err = gorm.Open(postgres.Open(str), &gorm.Config{})
		databaseSql, errSql = sql.Open("postgres", str)
	}

	if errSql != nil {
		log.Fatal("errorSql: ", errSql)
	}

	if err != nil {
		log.Fatal("error: ", err)
	}

	db = database
	dbSql = databaseSql

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}

func GetDatabaseSql() *sql.DB {
	return dbSql
}
