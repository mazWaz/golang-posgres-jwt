package user

import (
	"gorm.io/gorm"
)

type Role string

// Declare possible role values
const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

type ModelUser struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Role     Role   `gorm:"type:varchar(10);not null" json:"-"`
}

type RequestCreateUser struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,oneof=admin use"`
}

type RequestUpdateUser struct {
	Username string `json:"username" validate:"min=3,max=32"`
	Password string `json:"password" validate:"min=8"`
	Role     string `json:"role" validate:"oneof=admin use"`
}
