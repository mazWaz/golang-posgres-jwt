package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-clean/db"
	"go-clean/modules/user"
	"os"
	"time"
)

func generateAccessToken(user *user.ModelUser, expire int64, types Type) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"type": "access",
		"exp":  expire,
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func generateRefreshToken(user *user.ModelUser, expire int64, types Type) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"type": types,
		"exp":  expire,
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func saveToken(userId uint, token string, Type Type, Expires time.Time) error {
	var dataToken ModelToken
	dataToken.Token = token
	dataToken.UserId = userId
	dataToken.Type = Type
	dataToken.Expires = Expires
	return db.Data.Create(dataToken).Error
}

func generateAuthToken(user *user.ModelUser) (*ResponseAuthToken, error) {
	accessTokenExpire := time.Now().Add(time.Minute * 15).Unix() // Access token valid for 15 minutes
	accessToken, accessTokenErr := generateAccessToken(user, accessTokenExpire, Access)
	refreshTokenExpire := time.Now().Add(time.Hour * 72).Unix() // Refresh token valid for 72 hours
	refreshToken, refreshTokenErr := generateRefreshToken(user, refreshTokenExpire, Refresh)

	err := saveToken(user.ID, refreshToken, Refresh, time.Unix(refreshTokenExpire, 0))

	if err != nil {
		return nil, fmt.Errorf("fail save token")
	}

	if accessTokenErr != nil || refreshTokenErr != nil {
		return nil, fmt.Errorf("fail generate token")
	}

	return &ResponseAuthToken{
		Access: ResponseToken{
			Token:      accessToken,
			ExpireTime: time.Unix(accessTokenExpire, 0),
		},
		Refresh: ResponseToken{
			Token:      refreshToken,
			ExpireTime: time.Unix(refreshTokenExpire, 0),
		},
	}, nil
}
