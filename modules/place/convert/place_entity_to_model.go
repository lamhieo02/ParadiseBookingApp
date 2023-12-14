package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/place/iomodel"
)

func ConvertPlaceEntityToGetModel(data *entities.Place) *iomodel.GetPlaceResp {
	return &iomodel.GetPlaceResp{
		ID:            data.Id,
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
		MaxGuest:      data.MaxGuest,
		Numbed:        data.NumBed,
	}
}
