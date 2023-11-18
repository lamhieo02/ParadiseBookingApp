package placestorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *placeStorage) ListPlaceByVendorID(ctx context.Context, vendorID int) ([]entities.Place, error) {
	db := s.db

	var data []entities.Place

	if err := db.Where("vendor_id = ?", vendorID).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
