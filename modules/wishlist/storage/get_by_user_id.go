package wishliststorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *wishListStorage) GetByUserID(ctx context.Context, userId int) ([]entities.WishList, error) {
	db := s.db

	var data []entities.WishList

	if err := db.Where("user_id = ?", userId).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
