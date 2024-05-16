package bookingguiderstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *bookingGuiderSto) ListByCondition(ctx context.Context, conditions []common.Condition) ([]*entities.BookingGuider, error) {
	var res []*entities.BookingGuider
	db := s.db
	db = db.Table(entities.BookingGuider{}.TableName())

	for _, v := range conditions {
		query := v.BuildQuery()
		db = db.Where(query+" ?", v.Value)
	}

	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
