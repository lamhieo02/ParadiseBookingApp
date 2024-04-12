package placestorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *placeStorage) GetPlaceByLatLng(ctx context.Context, lat, lng float64) ([]*entities.Place, error) {
	var places []*entities.Place
	err := s.db.Raw("call GetPlaceByLatLng(?, ?)", lat, lng).Scan(&places).Error
	if err != nil {
		return nil, err
	}

	return places, nil
}
