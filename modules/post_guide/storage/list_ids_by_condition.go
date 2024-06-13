package postguidestorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postGuideStorage) ListPostGuideIdsByCondition(ctx context.Context, limit int, condition map[string]interface{}) ([]int, error) {
	var ids []int

	if err := s.db.Model(&entities.PostGuide{}).Where(condition).Pluck("id", &ids).Limit(limit).Error; err != nil {
		return nil, err
	}

	return ids, nil
}
