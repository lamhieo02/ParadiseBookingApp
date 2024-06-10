package requestvendorstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	requestvendoriomodel "paradise-booking/modules/request_vendor/iomodel"
)

func (s *requestVendorSto) ListByFilter(ctx context.Context, paging *common.Paging, filter *requestvendoriomodel.Filter) ([]*entities.RequestVendor, error) {
	db := s.db
	var data []*entities.RequestVendor

	db = db.Table(entities.RequestVendor{}.TableName()).Order("id desc")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
