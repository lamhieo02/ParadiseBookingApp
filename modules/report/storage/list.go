package reportstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	reportiomodel "paradise-booking/modules/report/iomodel"
)

func (s *reportStorage) ListReport(ctx context.Context, paging *common.Paging, filter *reportiomodel.Filter) ([]*entities.Report, error) {
	db := s.db

	var data []*entities.Report

	db = db.Table(entities.Report{}.TableName()).Order("id desc")

	if filter.ObjectID != 0 {
		db = db.Where("object_id = ?", filter.ObjectID)
	}

	if filter.ObjectType != 0 {
		db = db.Where("object_type = ?", filter.ObjectType)
	}

	if filter.StatusID != 0 {
		db = db.Where("status_id = ?", filter.StatusID)
	}

	if filter.UserID != 0 {
		db = db.Where("user_id = ?", filter.UserID)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil

}
