package bookingratingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
)

func (s *bookingratingstorage) GetByVendorID(ctx context.Context, vendorID int, objectType int) ([]entities.BookingRating, error) {
	db := s.db

	var data []entities.BookingRating

	if objectType == constant.BookingRatingObjectTypePlace {
		if err := db.Raw("call GetCommentsAndRatingsPlaceByVendorId(?)", vendorID).Scan(&data).Error; err != nil {
			return nil, common.ErrorDB(err)
		}
	} else if objectType == constant.BookingRatingObjectTypeGuide {
		if err := db.Raw("call GetCommentsAndRatingsPostGuideByVendorId(?)", vendorID).Scan(&data).Error; err != nil {
			return nil, common.ErrorDB(err)
		}
	}

	return data, nil
}
