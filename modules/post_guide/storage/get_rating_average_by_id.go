package postguidestorage

import (
	"context"
	"paradise-booking/constant"
)

func (s *postGuideStorage) GetRatingAverageByPostGuideId(ctx context.Context, postGuideId int64) (*float64, error) {
	var ratingAverage *float64
	err := s.db.Raw("call GetAverageRatingByObjectId(?, ?)", postGuideId, constant.BookingRatingObjectTypeGuide).Scan(&ratingAverage).Error
	if err != nil {
		return nil, err
	}

	return ratingAverage, nil
}
