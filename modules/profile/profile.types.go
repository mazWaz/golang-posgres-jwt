package profile

import "gorm.io/gorm"

type ModelAddress struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	UserID   uint   `gorm:"not null" json:"user_id"`
	Address  string `gorm:"not null" json:"address"`
	RT       string `gorm:"not null" json:"rt"`
	RW       string `gorm:"not null" json:"rw"`
	District string `gorm:"not null" json:"district"`
	City     string `gorm:"not null" json:"city"`
	Province string `gorm:"not null" json:"province"`
}

type RequestCreateAddress struct {
	Address  string `json:"address" validate:"required,min=3"`
	RT       string `json:"rt" validate:"required,gte=0"`
	RW       string `json:"rw" validate:"required,gte=0"`
	District string `json:"district" validate:"required,min=2"`
	City     string `json:"city" validate:"required,min=2"`
	Province string `json:"province" validate:"required,min=2"`
}

type RequestUpdateAddress struct {
	Address  string `json:"address" validate:"omitempty,min=3"`
	RT       string `json:"rt" validate:"omitempty,min=2"`
	RW       string `json:"rw" validate:"omitempty,min=2"`
	District string `json:"district" validate:"omitempty,min=2"`
	City     string `json:"city" validate:"omitempty,min=2"`
	Province string `json:"province" validate:"omitempty,min=2"`
}

type RequestQueryAddress struct {
	Username string `form:"username" validate:"omitempty"`
	Role     string `form:"role" validate:"omitempty,oneof=SUPERADMIN ADMIN USER"`
	Limit    int    `form:"limit" validate:"gte=1,omitempty,lte=100"`
	Page     int    `form:"page" validate:"gte=1,omitempty,lte=100"`
}

func (ModelAddress) TableName() string {
	return "user_address"
}
