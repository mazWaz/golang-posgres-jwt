package profile

import (
	"go-clean/db"
	"go-clean/utils"
)

type NewAddressService struct{}

func (s *NewAddressService) GetAddressByID(id uint) (*ModelAddress, error) {
	var address ModelAddress
	err := db.Data.First(&address, id).Error

	return &address, err
}

func (s *NewAddressService) GetAddressByUserID(id uint) (*ModelAddress, error) {
	var address ModelAddress
	err := db.Data.Joins("INNER JOIN users ON users.id = user_address.user_id").
		Where("users.id = ?", id).
		First(&address).Error
	return &address, err
}

func (s *NewAddressService) CreateAddress(address RequestCreateAddress, id float64) (ModelAddress, error) {
	// Assign form data JSON to struct
	var addressData ModelAddress

	addressData.UserID = uint(id)
	addressData.Address = address.Address
	addressData.RT = address.RT
	addressData.RW = address.RW
	addressData.District = address.District
	addressData.City = address.City
	addressData.Province = address.Province

	err := db.Data.Table("user_address").Create(address)

	if err != nil {

		return addressData, utils.SanitizeDBError(err.Error)

	}

	return addressData, nil
}

func (s *NewAddressService) UpdateAddress(id uint, input RequestUpdateAddress) error {
	address, err := Service.GetAddressByID(id)
	if err != nil {
		return utils.SanitizeDBError(err)
	}
	return utils.SanitizeDBError(db.Data.Model(&address).Updates(input).Error)
}

func (s *NewAddressService) DeleteAddress(id uint) error {
	var address ModelAddress
	return utils.SanitizeDBError(db.Data.Delete(&address, id).Error)
}

var Service = &NewAddressService{}
