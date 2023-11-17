package accountstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *accountStorage) GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error) {
	db := s.db
	var account entities.Account
	if err := db.Table(account.TableName()).Where("email = ?", email).First(&account).Error; err != nil {
		return nil, common.ErrorDB(err)
	}
	return &account, nil
}
