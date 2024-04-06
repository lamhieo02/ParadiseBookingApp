package postreviewstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postReviewStorage) UpdateWithMap(ctx context.Context, data *entities.PostReview, props map[string]interface{}) error {
	if err := s.db.Model(data).Updates(props).Error; err != nil {
		return err
	}

	return nil
}
