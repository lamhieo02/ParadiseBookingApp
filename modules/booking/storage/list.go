package bookingstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
	"time"
)

func (s *bookingStorage) ListByFilter(ctx context.Context, filter *iomodel.FilterListBooking, paging *common.Paging, userId int) ([]entities.Booking, error) {
	db := s.db

	var data []entities.Booking

	db = db.Table(entities.Booking{}.TableName())

	db = db.Where("user_id = ?", userId)
	if v := filter; v != nil {
		if len(v.Statuses) > 0 {
			db = db.Where("status_id in (?) ", v.Statuses)
		}

		if v.DateFrom != "" {
			dateTime := v.DateFrom + " 00:00:00"
			timeFrom, err := time.Parse("2006-01-02 15:04:05", dateTime)
			if err != nil {
				return nil, err
			}
			db = db.Where("created_at >= ?", timeFrom)
		}
		if v.DateTo != "" {
			dateTime := v.DateTo + " 00:00:00"
			timeTo, err := time.Parse("2006-01-02 15:04:05", dateTime)
			if err != nil {
				return nil, err
			}
			db = db.Where("created_at <= ?", timeTo)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
