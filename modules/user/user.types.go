package user

import (
	"go-clean/middlewares"
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

type ModelUserAddress struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	UserID   uint   `gorm:"not null" json:"user_id"`
	Address  string `gorm:"not null" json:"address"`
	RT       string `gorm:"not null" json:"-"`
	RW       string `gorm:"not null" json:"-"`
	District string `gorm:"not null" json:"-"`
	City     string `gorm:"not null" json:"-"`
	Province string `gorm:"not null" json:"-"`
}

func (ModelUserAddress) TableName() string {
	return "user_address"
}

func (ModelUser) TableName() string {
	return "users"
}

var ValidateQueryUser = middlewares.Validator{
	Query: &RequestQueryUser{},
}

var ValidateCreateUser = middlewares.Validator{
	Body: &RequestCreateUser{},
}

var ValidateUpdateUser = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
	Body: &RequestUpdateUser{},
}

var ValidateDeleteUser = middlewares.Validator{
	Param: &struct {
		Id string `uri:"id" validate:"required,gte=0"`
	}{},
}

type RequestQueryUser struct {
	Username string `form:"username" validate:"omitempty"`
	Role     string `form:"role" validate:"omitempty,oneof=SUPERADMIN ADMIN USER"`
	Limit    int    `form:"limit" validate:"gte=1,omitempty,lte=100"`
	Page     int    `form:"page" validate:"gte=1,omitempty,lte=100"`
}

type RequestCreateUser struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8"`
	Role     ROLE   `json:"role" validate:"oneof=SUPERADMIN ADMIN USER"`
}

type RequestUpdateUser struct {
	Username string `json:"username" validate:"omitempty,min=3,max=32"`
	Password string `json:"password" validate:"omitempty,min=8"`
	Role     string `json:"role" validate:"omitempty,oneof=SUPERADMIN ADMIN USER"`
}
