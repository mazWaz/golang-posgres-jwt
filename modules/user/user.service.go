package user

import (
	"go-clean/db"
	"go-clean/utils"
)

type NewUserService struct{}

func (s *NewUserService) GetUserByUsername(username string) (*ModelUser, error) {
	var user ModelUser

	if err := db.Data.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, utils.SanitizeDBError(err)
	}

	return &user, nil
}

func (s *NewUserService) GetUserByEmail(email string) (*ModelUser, error) {
	var user ModelUser
	err := db.Data.Where("email = ?", email).First(&user).Error
	return &user, utils.SanitizeDBError(err)
}

func (s *NewUserService) GetUserByID(id uint) (*ModelUser, error) {
	var user ModelUser
	err := db.Data.First(&user, id).Error

	return &user, err
}

func (s *NewUserService) GetAllUsers(filters map[string]interface{}) ([]ModelUser, error) {
	var users []ModelUser
	query := db.Data.Model(&[]ModelUser{})

	for key, value := range filters {
		if value != "" && value != "%%" { // Exclude empty and unmodified LIKE pattern
			query = query.Where(key, value)
		}
	}

	err := db.Data.Find(&users).Error
	return users, utils.SanitizeDBError(err)
}

func (s *NewUserService) CreateUser(user *ModelUser) error {
	return utils.SanitizeDBError(db.Data.Create(user).Error)
}

func (s *NewUserService) UpdateUser(id uint, input RequestUpdateUser) error {
	user, err := Service.GetUserByID(id)
	if err != nil {
		return utils.SanitizeDBError(err)
	}
	return utils.SanitizeDBError(db.Data.Model(&user).Updates(input).Error)
}

func (s *NewUserService) DeleteUser(id uint) error {
	var user ModelUser
	return utils.SanitizeDBError(db.Data.Delete(&user, id).Error)
}

var Service = &NewUserService{}
