package auth

import (
	"errors"
	"go-clean/db"
	"go-clean/modules/user"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type NewTokenService struct{}

func (s *NewTokenService) VerifyToken(refreshToken string, tokenType TokenType) (*ModelToken, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("INVALID Refresh Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("INVALID Refresh Token")
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
		return nil, errors.New("INVALID Refresh Token")
	}

	return &modelToken, nil
}

func (s *NewTokenService) GenerateToken(user *user.ModelUser) (*ResponseAuthToken, error) {
	accessTokenExpire := time.Now().Add(time.Minute * 15).Unix() // Access token valid for 15 minutes
	accessToken, accessTokenErr := TokenService.GenerateAccessToken(user, accessTokenExpire, Access)
	refreshTokenExpire := time.Now().Add(time.Hour * 72).Unix() // Refresh token valid for 72 hours
	refreshToken, refreshTokenErr := TokenService.GenerateRefreshToken(user, refreshTokenExpire, Refresh)

	if accessTokenErr != nil || refreshTokenErr != nil {
		return nil, errors.New("FAIL Generate Token")
	}

	deleteTokenErr := TokenService.DeleteRefreshToken(user.ID)

	if deleteTokenErr != nil {
		return nil, errors.New("FAIL Save Token")
	}

	saveTokenErr := TokenService.SaveToken(user.ID, refreshToken, Refresh, time.Unix(refreshTokenExpire, 0))

	if saveTokenErr != nil {
		return nil, errors.New("FAIL Save Token")
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
	dataToken.Type = Type
	dataToken.Expires = Expires
	dataToken.Blacklisted = false
	dataToken.UserID = userId

	return db.Data.Create(&dataToken).Error
}

func (s *NewTokenService) DeleteRefreshToken(userId uint) error {
	//TODO: Fix The Logic
	//Hard Delete
	return db.Data.Unscoped().Delete(&ModelToken{}, "user_id = ?", userId).Error
}

var TokenService = &NewTokenService{}
