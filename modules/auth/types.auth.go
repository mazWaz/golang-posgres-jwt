package auth

import (
	"go-clean/modules/user"
	"time"
)

type ROLE string

const (
	SUPERADMIN ROLE = "SUPERADMIN"
	ADMIN      ROLE = "ADMIN"
	USER       ROLE = "USER"
)

type TokenType string

const (
	Access        TokenType = "ACCESS"
	Refresh       TokenType = "REFRESH"
	ResetPassword TokenType = "RESET_PASSWORD"
	VerifyEmail   TokenType = "VERIFY_EMAIL"
)

type ModelToken struct {
	ID          uint           `gorm:"primaryKey;autoIncrement"`
	Token       string         `gorm:"unique;not null"`
	Type        TokenType      `gorm:"type:token_type;not null"`
	Expires     time.Time      `gorm:"not null"`
	Blacklisted bool           `gorm:"-" json:"-"`
	UserID      uint           `gorm:"not null"`
	User        user.ModelUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type RequestLogin struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8"`
}

type RequestRefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required,min=8"`
}

type ResponseToken struct {
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}

type ResponseAuthToken struct {
	Access  ResponseToken `json:"access"`
	Refresh ResponseToken `json:"refresh"`
}
