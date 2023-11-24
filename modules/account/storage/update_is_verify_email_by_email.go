package accountstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *accountStorage) UpdateIsVerifyEmailByEmail(ctx context.Context, email string) error {
	db := s.db
	account := entities.Account{
		IsEmailVerified: 1,
	}
	if err := db.Table(account.TableName()).Where("email = ?", email).Updates(&account).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
