package auth

import (
	"go-clean/modules/user"
	"gorm.io/gorm"
	"time"
)

type Type string

const (
	Access        Type = "ACCESS"
	Refresh       Type = "REFRESH"
	ResetPassword Type = "RESET_PASSWORD"
	VerifyEmail   Type = "VERIFY_EMAIL"
)

type ModelToken struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Token       string `gorm:"uniqueIndex;not null" json:"-"`
	Type        Type   `gorm:"type:varchar(10);not null" json:"-"`
	Expires     time.Time
	Blacklisted bool `gorm:"-" json:"-"`
	CreatedAt   time.Time
	UserId      uint
	User        user.ModelUser
}

type ValidateLogin struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8"`
}

type ValidateRefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required,min=8"`
}

type ResponsesToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserData     *user.ModelUser
}
