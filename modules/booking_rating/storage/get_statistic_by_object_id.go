package bookingratingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *bookingratingstorage) GetStatisticByObjectID(ctx context.Context, objectId int64, objectType int) ([]entities.StatisticResp, error) {
	db := s.db

	var data []entities.StatisticResp

	if err := db.Raw("call GetRatingStatisticByObjectId(?,?)", objectId, objectType).Scan(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
