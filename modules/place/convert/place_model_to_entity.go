package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/place/iomodel"
)

func ConvertPlaceCreateModelToEntity(data *iomodel.CreatePlaceReq) *entities.Place {
	return &entities.Place{
		Name:          data.Name,
		Description:   data.Description,
		PricePerNight: data.PricePerNight,
		Address:       data.Address,
		Capacity:      data.Capacity,
		Cover:         data.Cover,
		Lat:           data.Lat,
		Lng:           data.Lng,
		Country:       data.Country,
		State:         data.State,
		City:          data.City,
	}
}

func ConvertPlaceUpdateModelToEntity(data *iomodel.UpdatePlaceReq) *entities.Place {
	return &entities.Place{
		Name:          data.Name,
		Description:   data.Description,
		PricePerNight: data.PricePerNight,
		Address:       data.Address,
		Capacity:      data.Capacity,
		Cover:         data.Cover,
		Lat:           data.Lat,
		Lng:           data.Lng,
		Country:       data.Country,
		State:         data.State,
		City:          data.City,
	}
}