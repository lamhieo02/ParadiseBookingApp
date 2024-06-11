package requestguiderstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	requestguideriomodel "paradise-booking/modules/request_guider/iomodel"
)

func (s *RequestGuiderSto) ListByFilter(ctx context.Context, paging *common.Paging, filter *requestguideriomodel.Filter) ([]*entities.RequestGuider, error) {
	db := s.db
	var data []*entities.RequestGuider

	db = db.Table(entities.RequestGuider{}.TableName()).Order("id desc")

	if filter.UserID != 0 {
		db = db.Where("user_id = ?", filter.UserID)
	}

	if filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
