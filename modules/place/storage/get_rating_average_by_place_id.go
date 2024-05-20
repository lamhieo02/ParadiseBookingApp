package placestorage

import (
	"context"
	"paradise-booking/constant"
)

func (s *placeStorage) GetRatingAverageByPlaceId(ctx context.Context, placeId int64) (*float64, error) {
	var ratingAverage *float64
	err := s.db.Raw("call GetAverageRatingByObjectId(?, ?)", placeId, constant.BookingRatingObjectTypePlace).Scan(&ratingAverage).Error
	if err != nil {
		return nil, err
	}

	return ratingAverage, nil
}
