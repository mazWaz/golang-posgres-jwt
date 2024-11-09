package migrations

import (
	"go-clean/modules/auth"
	"go-clean/modules/user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := CreateTokenTypeEnum(db); err != nil {
		if err.Error() != "pq: type \"token_type\" already exists" {
			return err
		}
	}

	if err := CreateRoleTypeEnum(db); err != nil {
		if err.Error() != "pq: type \"role_type\" already exists" {
			return err
		}
	}

	if err := db.AutoMigrate(
		&auth.ModelToken{},
		&user.ModelUser{},
	); err != nil {
		return err
	}

	return nil
}

func CreateTokenTypeEnum(db *gorm.DB) error {
	db.Exec("DROP TYPE IF EXISTS token_type CASCADE")
	return db.Exec("CREATE TYPE token_type AS ENUM ('ACCESS', 'REFRESH', 'RESET_PASSWORD', 'VERIFY_EMAIL')").Error

}

func CreateRoleTypeEnum(db *gorm.DB) error {
	db.Exec("DROP TYPE IF EXISTS role_type CASCADE")
	return db.Exec("CREATE TYPE role_type AS ENUM ('SUPERADMIN', 'ADMIN', 'USER')").Error
}
