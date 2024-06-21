package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/place/iomodel"
	"strings"
)

func ConvertPlaceEntityToGetModel(data *entities.Place, isFree bool, ratingAverage *float64) *iomodel.GetPlaceResp {
	res := &iomodel.GetPlaceResp{
		ID:            data.Id,
		VendorID:      data.VendorID,
		Name:          data.Name,
		Description:   data.Description,
		PricePerNight: data.PricePerNight,
		Address:       data.Address,
		Images:        []string{},
		Lat:           data.Lat,
		Lng:           data.Lng,
		Country:       data.Country,
		State:         data.State,
		District:      data.District,
		MaxGuest:      data.MaxGuest,
		Numbed:        data.NumBed,
		IsFree:        isFree,
		RatingAverage: *ratingAverage,
		BedRoom:       data.BedRoom,
	}

	if data.Cover != "" {
		res.Images = strings.Split(data.Cover, ",")
	}

	return res
}
