package main

import (
	"log"

	. "iabots-server/configs"
	. "iabots-server/internal/infra/database"
	. "iabots-server/internal/infra/database/seeds"
)

func main() {
	log.Println("🌱 Running database seeds")

	LoadEnv()

	db, err := NewDatabase()
	if err != nil {
		log.Fatalf("database error: %v", err)
	}

	if err := SeedModelProviders(db.GetDB()); err != nil {
		log.Fatalf("seed error: %v", err)
	}

	log.Println("✅ Seeds finished successfully")
}
