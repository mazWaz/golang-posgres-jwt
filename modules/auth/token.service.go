package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-clean/db"
	"go-clean/modules/user"
	"os"
	"time"
)

type NewTokenService struct{}

func (s *NewTokenService) VerifyToken(refreshToken string, tokenType TokenType) (*ModelToken, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("INVALID Refresh Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("INVALID Refresh Token")
	}

	userId := claims["sub"].(string)

	var modelToken ModelToken
	tokenData := db.Data.Where(
		"token = ? AND "+
			"userId = ? AND "+
			"type = ? AND "+
			"blacklisted = FALSE",
		refreshToken,
		userId,
		tokenType).First(&modelToken)

	if tokenData.Error != nil {
		return nil, fmt.Errorf("FAIL Token Not Found")
	}

	return &modelToken, nil

}

func (s *NewTokenService) GenerateAccessToken(user *user.ModelUser, expire int64, types TokenType) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"type": types,
		"exp":  expire,
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s *NewTokenService) GenerateRefreshToken(user *user.ModelUser, expire int64, types TokenType) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"type": types,
		"exp":  expire,
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s *NewTokenService) SaveToken(userId uint, token string, Type TokenType, Expires time.Time) error {
	var dataToken ModelToken
	dataToken.Token = token
	dataToken.UserID = userId
	dataToken.Type = Type
	dataToken.Expires = Expires
	return db.Data.Create(dataToken).Error
}

func (s *NewTokenService) GenerateToken(user *user.ModelUser) (*ResponseAuthToken, error) {
	accessTokenExpire := time.Now().Add(time.Minute * 15).Unix() // Access token valid for 15 minutes
	accessToken, accessTokenErr := TokenService.GenerateAccessToken(user, accessTokenExpire, Access)
	refreshTokenExpire := time.Now().Add(time.Hour * 72).Unix() // Refresh token valid for 72 hours
	refreshToken, refreshTokenErr := TokenService.GenerateRefreshToken(user, refreshTokenExpire, Refresh)

	err := TokenService.SaveToken(user.ID, refreshToken, Refresh, time.Unix(refreshTokenExpire, 0))

	if err != nil {
		return nil, fmt.Errorf("FAIL Save Token")
	}

	if accessTokenErr != nil || refreshTokenErr != nil {
		return nil, fmt.Errorf("FAIL Generate Token")
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

var TokenService = &NewTokenService{}
