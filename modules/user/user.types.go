package user

import (
	"gorm.io/gorm"
)

type ROLE string

const (
	SUPERADMIN ROLE = "SUPERADMIN"
	ADMIN      ROLE = "ADMIN"
	USER       ROLE = "USER"
)

type ModelUser struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Role     ROLE   `gorm:"type:role_type;not null"`
}

func (ModelUser) TableName() string {
	return "users"
}

type RequestQueryUser struct {
	Username string `json:"username" validate:"min=3,max=32"`
	Password string `json:"password" validate:"min=8"`
	Role     string `json:"role" validate:"oneof=SUPERADMIN ADMIN USER"`
	Limit    uint8  `json:"limit" validate:"gte=0,lte=100"`
	Page     uint8  `json:"page" validate:"gte=0,lte=100"`
}

type RequestCreateUser struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"oneof=SUPERADMIN ADMIN USER"`
}

type RequestUpdateUser struct {
	Username string `json:"username" validate:"min=3,max=32"`
	Password string `json:"password" validate:"min=8"`
	Role     string `json:"role" validate:"oneof=admin use"`
}
