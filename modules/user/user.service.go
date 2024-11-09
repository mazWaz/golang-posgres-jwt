package user

import (
	"fmt"
	"go-clean/db"
	"go-clean/middlewares"
	"go-clean/utils"

	"gorm.io/gorm"
)

type NewUserService struct{}

func (s *NewUserService) GetUserByUsernameOrEmail(credential string) (*ModelUser, error) {
	var user ModelUser

	if err := db.Data.Where("username = ?", credential).Or("email = ?", credential).Find(&user).Error; err != nil {
		return nil, utils.SanitizeDBError(err)
	}

	return &user, nil
}

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

func (s *NewUserService) CreateEmailByAdmin(req *RequestCreateUserByAdmin) (ModelUser, error) {
	var userData ModelUser

	// Assign form data JSON to struct
	userData.Email = &req.Email
	userData.Role = req.Role

	return userData, utils.SanitizeDBError(db.Data.Create(&userData).Error)
}

func (s *NewUserService) CreateByAdmin(req *RequestCreateUserByAdmin) (*UserProfile, error) {
	// Assign form data JSON to struct
	var addressData ModelAddress
	var userData ModelUser

	userData.Username = &req.Username
	if req.Password != "" {
		hashedPassword, _ := utils.HashPassword([]byte(req.Password))
		userData.Password = string(hashedPassword)
	}
	userData.Email = &req.Email

	addressData.Address = req.Address
	addressData.RT = int(req.RT)
	addressData.RW = int(req.RW)
	addressData.District = req.District
	addressData.City = req.City
	addressData.Province = req.Province

	// DB Transaction
	err := db.Data.Transaction(func(tx *gorm.DB) error {
		_ = tx.Create(&userData)
		addressData.UserID = userData.ID

		if err := tx.Debug().Create(&addressData).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, utils.SanitizeDBError(err)
	}

	// Return
	return &UserProfile{
		userData.ID,
		userData.Username,
		userData.Email,
		addressData,
	}, nil
}

func (s *NewUserService) CreateUser(role string, req *RequestCreateUser) (*UserProfile, error) {
	var addressData ModelAddress
	var userData ModelUser

	// Conditional check on unique type field and password
	if req.Username != nil && *req.Username != "" {
		userData.Username = req.Username
	}

	if req.Password != "" {
		hashedPassword, _ := utils.HashPassword([]byte(req.Password))
		userData.Password = string(hashedPassword)
	}

	if role == string(middlewares.ADMIN) && req.Role == middlewares.SUPERADMIN {
		return nil, fmt.Errorf("user is unauthorized to add SUPERADMIN role")
	}

	// Assign form data JSON to struct
	userData.Email = req.Email
	userData.Role = req.Role

	addressData.Address = req.Address
	addressData.RT = int(req.RT)
	addressData.RW = int(req.RW)
	addressData.District = req.District
	addressData.City = req.City
	addressData.Province = req.Province

	// DB Transaction
	err := db.Data.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Create(&userData).Error; err != nil {
			return err
		}
		addressData.UserID = userData.ID

		if err := tx.Debug().Create(&addressData).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, utils.SanitizeDBError(err)
	}

	return &UserProfile{
		userData.ID,
		userData.Username,
		userData.Email,
		addressData,
	}, nil
}

func (s *NewUserService) UpdateAddress(userId uint, role string, req *RequestUpdateAddress) (*UserProfile, error) {
	var addressData ModelAddress
	var userData ModelUser

	// Conditional check on unique type field and password
	if req.Username != nil && *req.Username != "" {
		userData.Username = req.Username
	}

	if req.Password != "" {
		hashedPassword, _ := utils.HashPassword([]byte(req.Password))
		userData.Password = string(hashedPassword)
	}

	updatedData, _ := Service.GetFullUserByUserID(uint(userId))

	addressData.ID = updatedData.ID

	// Assign form data JSON to struct
	userData.ID = uint(userId)
	userData.Role = middlewares.ROLE(role)
	userData.Email = req.Email
	addressData.Address = req.Address
	addressData.RT = int(req.RT)
	addressData.RW = int(req.RW)
	addressData.District = req.District
	addressData.City = req.City
	addressData.Province = req.Province

	// DB Transaction
	err := db.Data.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Where("user_id = ?", userId).Model(&addressData).Updates(addressData).Error; err != nil {
			return err
		}

		if err := tx.Debug().Model(&userData).Updates(userData).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, utils.SanitizeDBError(err)
	}

	return &UserProfile{
		userData.ID,
		userData.Username,
		userData.Email,
		addressData,
	}, nil
}

func (s *NewUserService) GetFullUserByUserID(id uint) (*UserProfile, error) {
	var addressData ModelAddress
	var userData ModelUser

	user, errUser := Service.GetUserByID(id)

	if errUser != nil {
		return nil, utils.SanitizeDBError(errUser)
	}

	userData.ID = user.ID
	userData.Username = user.Username
	userData.Email = user.Email

	errAddress := db.Data.Debug().Joins("INNER JOIN users ON users.id = user_address.user_id").
		Where("users.id = ?", id).
		First(&addressData).Error

	return &UserProfile{
		userData.ID,
		userData.Username,
		userData.Email,
		addressData,
	}, utils.SanitizeDBError(errAddress)
}

func (s *NewUserService) UpdateUser(role string, UserId uint, req *RequestUpdateUser) (*UserProfile, error) {
	// Assign form data JSON to struct
	var addressData ModelAddress
	var userData ModelUser

	// Conditional check on unique type field and password
	if req.Username != nil && *req.Username != "" {
		userData.Username = req.Username
	}

	if req.Password != "" {
		hashedPassword, _ := utils.HashPassword([]byte(req.Password))
		userData.Password = string(hashedPassword)
	}

	user, _ := Service.GetUserByID(UserId)

	if role == string(middlewares.ADMIN) && user.Role == middlewares.SUPERADMIN {
		return nil, fmt.Errorf("user is unauthorized to change SUPERADMIN role")
	}

	updatedData, _ := Service.GetFullUserByUserID(uint(UserId))

	addressData.ID = updatedData.ID
	addressData.Address = req.Address
	addressData.RT = int(req.RT)
	addressData.RW = int(req.RW)
	addressData.District = req.District
	addressData.City = req.City
	addressData.Province = req.Province

	userData.ID = uint(UserId)
	userData.Username = req.Username
	userData.Email = req.Email

	// DB Transaction
	errTransaction := db.Data.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Where("user_id = ?", UserId).Model(&addressData).Updates(addressData).Error; err != nil {
			return err
		}

		if err := tx.Debug().Model(&userData).Updates(userData).Error; err != nil {
			return err
		}
		return nil
	})

	if errTransaction != nil {
		return nil, utils.SanitizeDBError(errTransaction)
	}

	return &UserProfile{
		userData.ID,
		userData.Username,
		userData.Email,
		addressData,
	}, nil
}

func (s *NewUserService) DeleteUser(role string, userId uint) error {
	var addressData ModelAddress
	var userData ModelUser

	user, _ := Service.GetUserByID(userId)

	if role == string(middlewares.ADMIN) && user.Role == middlewares.SUPERADMIN {
		return fmt.Errorf("user is unauthorized to change SUPERADMIN role")
	}

	// DB Transaction
	errTransaction := db.Data.Transaction(func(tx *gorm.DB) error {

		if err := tx.Debug().Where("user_id = ?", userId).Model(&addressData).Delete(&addressData).Error; err != nil {
			return err
		}

		if err := tx.Debug().Where("id = ?", userId).Model(&userData).Delete(&userData).Error; err != nil {
			return err
		}
		return nil
	})

	return utils.SanitizeDBError(errTransaction)
}

var Service = &NewUserService{}
