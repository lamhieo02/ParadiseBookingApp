package accountstorage

import (
	"context"
	"paradise-booking/constant"
	"paradise-booking/entities"
)

func (s *accountStorage) GetAllAccountUserAndVendor(ctx context.Context) ([]entities.Account, error) {
	var result []entities.Account
	db := s.db.Table(entities.Account{}.TableName())
	err := db.Where("role = ? OR role = ?", constant.UserRole, constant.VendorRole).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil

}
