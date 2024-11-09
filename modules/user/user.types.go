package user

import (
	"go-clean/middlewares"

	"gorm.io/gorm"
)

type ModelUser struct {
	gorm.Model
	ID       uint             `gorm:"primaryKey;autoIncrement"`
	Username *string          `gorm:"uniqueIndex;not null" json:"username"`
	Email    *string          `gorm:"uniqueIndex;not null" json:"email"`
	Password string           `gorm:"not null" json:"-"`
	Role     middlewares.ROLE `gorm:"type:role_type;not null"`
}

type ModelAddress struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	UserID   uint   `gorm:"not null" json:"user_id"`
	Address  string `gorm:"not null" json:"address"`
	RT       int    `gorm:"not null" json:"rt"`
	RW       int    `gorm:"not null" json:"rw"`
	District string `gorm:"not null" json:"district"`
	City     string `gorm:"not null" json:"city"`
	Province string `gorm:"not null" json:"province"`
}

func (ModelUser) TableName() string {
	return "users"
}

func (ModelAddress) TableName() string {
	return "user_address"
}

type RequestQueryUserByAdmin struct {
	Username string `form:"username" validate:"omitempty"`
	Role     string `form:"role" validate:"omitempty,oneof=SUPERADMIN ADMIN USER"`
	Limit    int    `form:"limit" validate:"gte=1,omitempty,lte=100"`
	Page     int    `form:"page" validate:"gte=1,omitempty,lte=100"`
}

type RequestCreateUserByAdmin struct {
	Username string           `json:"username" validate:"omitempty"`
	Password string           `json:"password" validate:"omitempty,min=6"`
	Email    string           `json:"email" validate:"required,email,min=8"`
	Role     middlewares.ROLE `json:"role" validate:"oneof=SUPERADMIN ADMIN USER,required"`
	Address  string           `json:"address" validate:"omitempty,min=3"`
	RT       int              `json:"rt" validate:"gte=0,omitempty"`
	RW       int              `json:"rw" validate:"gte=0,omitempty"`
	District string           `json:"district" validate:"omitempty,min=2"`
	City     string           `json:"city" validate:"omitempty,min=2"`
	Province string           `json:"province" validate:"omitempty,min=2"`
}

type RequestCreateUser struct {
	Username *string          `json:"username" validate:"omitempty"`
	Password string           `json:"password" validate:"omitempty"`
	Email    *string          `json:"email" validate:"required,email,min=6"`
	Role     middlewares.ROLE `json:"role" validate:"oneof=SUPERADMIN ADMIN USER,required"`
	Address  string           `json:"address" validate:"omitempty,min=5"`
	RT       int              `json:"rt" validate:"gte=0,omitempty"`
	RW       int              `json:"rw" validate:"gte=0,omitempty"`
	District string           `json:"district" validate:"omitempty,min=2"`
	City     string           `json:"city" validate:"omitempty,min=2"`
	Province string           `json:"province" validate:"omitempty,min=2"`
}

type RequestCreateEmailByAdmin struct {
	Email string           `json:"email" validate:"required,email,min=8"`
	Role  middlewares.ROLE `json:"role" validate:"oneof=SUPERADMIN ADMIN USER,required"`
}

type RequestUpdateUserByAdmin struct {
	Username string           `json:"username" validate:"omitempty"`
	Password string           `json:"password" validate:"omitempty,min=8"`
	Email    string           `json:"email" validate:"omitempty,email"`
	Role     middlewares.ROLE `json:"role" validate:"omitempty,oneof=SUPERADMIN ADMIN USER"`
	Address  string           `json:"address" validate:"omitempty,min=3"`
	RT       int              `json:"rt" validate:"gte=0,omitempty"`
	RW       int              `json:"rw" validate:"gte=0,omitempty"`
	District string           `json:"district" validate:"omitempty,min=2"`
	City     string           `json:"city" validate:"omitempty,min=2"`
	Province string           `json:"province" validate:"omitempty,min=2"`
}

type RequestUpdateAddress struct {
	Address  string  `json:"address" validate:"required,min=3"`
	RT       int     `json:"rt" validate:"gte=0,required"`
	RW       int     `json:"rw" validate:"gte=0,required"`
	District string  `json:"district" validate:"required,min=2"`
	City     string  `json:"city" validate:"required,min=2"`
	Province string  `json:"province" validate:"required,min=2"`
	Username *string `json:"username" validate:"omitempty"`
	Password string  `json:"password" validate:"omitempty,min=6"`
	Email    *string `json:"email" validate:"omitempty,min=6"`
}

type RequestUpdateUser struct {
	Username *string          `json:"username" validate:"omitempty"`
	Password string           `json:"password" validate:"omitempty,min=6"`
	Email    *string          `json:"email" validate:"omitempty,email"`
	Role     middlewares.ROLE `json:"role" validate:"omitempty,oneof=SUPERADMIN ADMIN USER"`
	Address  string           `json:"address" validate:"omitempty,min=3"`
	RT       int              `json:"rt" validate:"gte=0,omitempty"`
	RW       int              `json:"rw" validate:"gte=0,omitempty"`
	District string           `json:"district" validate:"omitempty,min=2"`
	City     string           `json:"city" validate:"omitempty,min=2"`
	Province string           `json:"province" validate:"omitempty,min=2"`
}

type UserProfile struct {
	UserId   uint    `json:"user_id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	ModelAddress
}
