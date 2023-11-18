package placestorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/place/iomodel"
)

func (s *placeStorage) ListPlaces(ctx context.Context, paging *common.Paging, filter *iomodel.Filter) ([]entities.Place, error) {
	db := s.db

	var data []entities.Place

	db = db.Table(entities.Place{}.TableName())

	if v := filter; v != nil {
		if v.VendorID != nil {
			db = db.Where("vendor_id = ?", v.VendorID)
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
