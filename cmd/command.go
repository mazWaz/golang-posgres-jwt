package cmd

import (
	"go-clean/db/migrations"
	"gorm.io/gorm"
	"log"
	"os"
)

func Commands(db *gorm.DB) {
	migrate := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
	}

	if migrate {
		if err := migrations.Migrate(db); err != nil {
			log.Fatalf("error migration: %v", err)
		}
		log.Println("migration completed successfully")
	}
}
