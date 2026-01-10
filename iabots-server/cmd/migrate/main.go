package main

import (
	"log"

	. "iabots-server/configs"
	. "iabots-server/internal/infra/database"
)

func main() {

	LoadEnv()

	db, err := NewDatabase()
	if err != nil {
		log.Fatalf("database connect error: %v", err)
	}

	if err := AutoMigrate(db.DB); err != nil {
		log.Fatalf("database migrate error: %v", err)
	}

	log.Println("✅ Database migrated successfully")
}
