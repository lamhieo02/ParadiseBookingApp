package calendarguiderstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
)

func (s *calendarGuiderStorage) ListByFilter(ctx context.Context, paging *common.Paging, filter *calendarguideriomodel.Filter) ([]*entities.CalendarGuider, error) {
	db := s.db.Table(entities.CalendarGuider{}.TableName()).Order("id DESC")
	var result []*entities.CalendarGuider
	if filter.DateFrom != "" {
		db = db.Where("date_from >= ?", filter.DateFrom)
	}

	if filter.DateTo != "" {
		db = db.Where("date_to <= ?", filter.DateTo)
	}

	if filter.GuiderID != 0 {
		db = db.Where("guider_id = ?", filter.GuiderID)
	}

	if filter.PostGuideID != 0 {
		db = db.Where("post_guide_id = ?", filter.PostGuideID)
	}

	if filter.PricePerPerson != 0 {
		db = db.Where("price_per_person = ?", filter.PricePerPerson)
	}

	if filter.Status != nil {
		db = db.Where("status = ?", filter.Status)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
