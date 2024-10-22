package migrations

import (
	"go-clean/modules/auth"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// Create the enum type in PostgreSQL if it does not exist
	db.Exec(`
        DO $$ 
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'TokenType') THEN
                CREATE TYPE "TokenType" AS ENUM ('ACCESS', 'REFRESH', 'RESET_PASSWORD', 'VERIFY_EMAIL');
            END IF;
        END $$;
    `)

	// Migrate the Token model
	return db.AutoMigrate(&auth.ModelToken{})
}
