package likepostreviewstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *likePostReviewStorage) UpdateWithMap(ctx context.Context, data *entities.LikePostReview, props map[string]interface{}) error {
	if err := s.db.Model(&data).Updates(props).Error; err != nil {
		return err
	}
	return nil
}
