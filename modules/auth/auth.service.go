package auth

import (
	"fmt"
	"go-clean/db"
	"go-clean/modules/user"
	"golang.org/x/crypto/bcrypt"
)

func LoginWithUsernameAndPassword(username string, password string) (*user.ModelUser, error) {

	userData, err := user.GetUserByUsername(username)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password)) != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	userData.Password = ""

	return userData, nil
}

func LogoutWithRefreshToken(refreshToken string) error {
	var modelToken ModelToken
	tokenData := db.Data.Where(
		"token = ? AND "+
			"type = ? AND "+
			"blacklisted = FALSE",
		refreshToken,
		Refresh,
	).First(&modelToken)

	if tokenData.Error != nil {
		return fmt.Errorf("FAIL Token Not Found")
	}
	if err := db.Data.Delete(&modelToken, modelToken.ID); err != nil {
		return fmt.Errorf("ERROR Deleting User")
	}
	return nil
}

func RefreshAuth(refreshToken string) (*ResponseAuthToken, error) {
	refreshTokenData, tokenError := VerifyToken(refreshToken, Refresh)
	if tokenError != nil {
		return nil, tokenError
	}

	userId := refreshTokenData.UserID

	userData, err := user.GetUserByID(userId)
	if err != nil {
		return nil, fmt.Errorf("ERROR User not Found")
	}

	var modelToken ModelToken
	if err := db.Data.Delete(&modelToken, refreshTokenData.ID); err != nil {
		return nil, fmt.Errorf("ERROR Deleting User")
	}

	token, tokenErr := GenerateToken(userData)
	if tokenErr != nil {
		return nil, fmt.Errorf("ERROR Generating Token")

	}

	return token, nil

}
