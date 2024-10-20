package auth

import (
	"fmt"
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
