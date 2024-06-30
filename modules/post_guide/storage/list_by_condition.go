package postguidestorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *postGuideStorage) ListPostGuideByCondition(ctx context.Context, conditions []common.Condition) ([]entities.PostGuide, error) {
	var data []entities.PostGuide

	db := s.db
	db = db.Table(entities.PostGuide{}.TableName())

	for _, v := range conditions {
		query := v.BuildQuery()
		db = db.Where(query+" ?", v.Value)
	}

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil

}
