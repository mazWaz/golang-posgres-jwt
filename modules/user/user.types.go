package user

import (
	"go-clean/middlewares"

	"gorm.io/gorm"
)

type ModelUser struct {
	gorm.Model
	ID       uint             `gorm:"primaryKey;autoIncrement"`
	Username string           `gorm:"uniqueIndex;not null" json:"username"`
	Email    string           `gorm:"not null" json:"email"`
	Password string           `gorm:"not null" json:"-"`
	Role     middlewares.ROLE `gorm:"type:role_type;not null"`
}

func (ModelUser) TableName() string {
	return "users"
}

type RequestQueryUser struct {
	Username string `form:"username" validate:"omitempty"`
	Role     string `form:"role" validate:"omitempty,oneof=SUPERADMIN ADMIN USER"`
	Limit    int    `form:"limit" validate:"gte=1,omitempty,lte=100"`
	Page     int    `form:"page" validate:"gte=1,omitempty,lte=100"`
}

type RequestCreateUserAdmin struct {
	Username string           `json:"username" validate:"required"`
	Password string           `json:"password" validate:"required,min=8"`
	Email    string           `json:"email" validate:"required,email,min=8"`
	Role     middlewares.ROLE `json:"role" validate:"oneof=SUPERADMIN ADMIN USER"`
}

type RequestCreateUser struct {
	Email string           `json:"email" validate:"required,email,min=8"`
	Role  middlewares.ROLE `json:"role" validate:"oneof=SUPERADMIN ADMIN USER"`
}

type RequestUpdateUser struct {
	Username string `json:"username" validate:"omitempty,min=3,max=32"`
	Password string `json:"password" validate:"omitempty,min=8"`
	Role     string `json:"role" validate:"omitempty,oneof=SUPERADMIN ADMIN USER"`
}
