package bookingguiderstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
	"paradise-booking/utils"
)

func (s *bookingGuiderSto) ListByFilter(ctx context.Context, paging *common.Paging, filter *bookingguideriomodel.Filter, userId int) ([]entities.BookingGuider, error) {
	db := s.db

	var data []entities.BookingGuider

	db = db.Table(entities.BookingGuider{}.TableName())

	db = db.Where("user_id = ?", userId)
	if v := filter; v != nil {
		if len(v.Statuses) > 0 && v.Statuses[0] != 0 {
			db = db.Where("status_id in (?) ", v.Statuses)
		}

		if v.DateFrom != "" {
			dateTime, _ := utils.ParseStringToTime(v.DateFrom)
			db = db.Where("created_at >= ?", dateTime)
		}
		if v.DateTo != "" {
			dateTime, _ := utils.ParseStringToTime(v.DateTo)
			db = db.Where("created_at <= ?", dateTime)
		}

		if v.PostGuideID != 0 {
			db = db.Where("post_guide_id = ?", v.PostGuideID)
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
