package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
)

func ConvertBookingModelToListBooking(data entities.Booking, place *entities.Place) iomodel.DataListBooking {
	dataBooking := iomodel.DataListBooking{
		Id:          data.Id,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		PlaceId:     data.PlaceId,
		Place:       *place,
		StatusId:    data.StatusId,
		CheckInDate: data.ChekoutDate,
		ChekoutDate: data.ChekoutDate,
	}
	return dataBooking
}
