package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
)

func ConvertBookingModelToBookingEntity(model *iomodel.CreateBookingReq) *entities.Booking {
	return &entities.Booking{
		UserId:      model.UserID,
		PlaceId:     model.PlaceID,
		CheckInDate: model.CheckInDate,
		ChekoutDate: model.CheckOutDate,
	}
}
