package placestorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *placeStorage) ListPlaceIdsByCondition(ctx context.Context, limit int, condition map[string]interface{}) ([]int, error) {
	var ids []int
	err := s.db.Model(&entities.Place{}).Where(condition).Pluck("id", &ids).Limit(limit).Error
	if err != nil {
		return nil, err
	}

	return ids, nil
}
