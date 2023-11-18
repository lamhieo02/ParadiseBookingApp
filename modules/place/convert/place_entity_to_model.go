package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/place/iomodel"
)

func ConvertPlaceEntityToGetModel(data *entities.Place) *iomodel.GetPlaceResp {
	return &iomodel.GetPlaceResp{
		VendorID:      data.VendorID,
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
