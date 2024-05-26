package amenitystorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *amenityStorage) ListByObjectID(ctx context.Context, objectID int, objectType int) ([]entities.Amenity, error) {
	var res []entities.Amenity
	db := s.db.Table(entities.Amenity{}.TableName())
	if err := db.Where("object_id = ? AND object_type = ?", objectID, objectType).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
