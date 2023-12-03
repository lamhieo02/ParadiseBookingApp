package bookingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"

	"gorm.io/gorm"
)

func (s *bookingStorage) GetByPlaceID(ctx context.Context, placeId int) (*entities.Booking, error) {
	db := s.db

	var data entities.Booking

	if err := db.Where("place_id = ?", placeId).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, common.ErrorDB(err)
	}

	return &data, nil
}
