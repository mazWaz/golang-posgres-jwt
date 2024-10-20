package user

import (
	"go-clean/db"
)

func GetUserByUsername(username string) (*ModelUser, error) {
	var user ModelUser
	if err := db.Data.Find("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *ModelUser) error {
	return db.Data.Create(user).Error
}

func GetUserByEmail(email string) (*ModelUser, error) {
	var user ModelUser
	err := db.Data.Where("email = ?", email).First(&user).Error
	return &user, err
}

func GetUserByID(id uint) (*ModelUser, error) {
	var user ModelUser
	err := db.Data.First(&user, id).Error
	return &user, err
}

func GetAllUsers() ([]ModelUser, error) {
	var users []ModelUser
	err := db.Data.Find(&users).Error
	return users, err
}

func UpdateUser(user *ModelUser) error {
	return db.Data.Save(user).Error
}

func DeleteUser(id uint) error {
	var user ModelUser
	return db.Data.Delete(&user, id).Error
}
