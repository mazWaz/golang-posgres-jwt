package auth

import (
	"errors"
	"go-clean/db"
	"go-clean/modules/user"
	"go-clean/utils"
)

type NewAuthService struct{}

func (s *NewAuthService) LoginWithUsernameAndEmail(credential string, password string) (*user.ModelUser, error) {

	userData, err := user.Service.GetUserByUsernameOrEmail(credential)
	if err != nil || userData == nil {
		return nil, errors.New("invalid credentials")
	}

	hashErr := utils.ComparePassword(password, userData.Password)
	if !hashErr {
		return nil, errors.New("invalid credentials")
	}

	return userData, nil
}

func (s *NewAuthService) LoginWithUsernameAndPassword(username string, password string) (*user.ModelUser, error) {

	userData, err := user.Service.GetUserByUsername(username)
	if err != nil || userData == nil {
		return nil, errors.New("invalid Username or Password")
	}

	hashErr := utils.ComparePassword(password, userData.Password)
	if !hashErr {
		return nil, errors.New("invalid Username or Password")
	}

	return userData, nil
}

func (s *NewAuthService) LogoutWithRefreshToken(refreshToken string) error {
	var modelToken ModelToken
	tokenData := db.Data.Where(
		"token = ? AND "+
			"type = ? AND "+
			"blacklisted = FALSE",
		refreshToken,
		Refresh,
	).First(&modelToken)

	if tokenData.Error != nil {
		return errors.New("FAIL Token Not Found")
	}
	if err := db.Data.Delete(&modelToken, modelToken.ID); err != nil {
		return errors.New("ERROR Deleting User")
	}
	return nil
}

func (s *NewAuthService) RefreshAuth(refreshToken string) (*ResponseAuthToken, error) {
	refreshTokenData, tokenError := TokenService.VerifyToken(refreshToken, Refresh)
	if tokenError != nil {
		return nil, tokenError
	}

	userId := refreshTokenData.UserID

	userData, err := user.Service.GetUserByID(userId)
	if err != nil {
		return nil, errors.New("ERROR User not Found")
	}

	token, tokenErr := TokenService.GenerateToken(userData)
	if tokenErr != nil {
		return nil, errors.New("ERROR Generating Token")

	}

	return token, nil

}

var AuthService = &NewAuthService{}
