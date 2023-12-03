package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
)

func ConvertBookingModelToGetByPlaceResp(user *entities.Account, dataBooking *entities.Booking, place *entities.Place) *iomodel.GetBookingByPlaceResp {
	return &iomodel.GetBookingByPlaceResp{
		Id:          dataBooking.Id,
		CreatedAt:   dataBooking.CreatedAt,
		UpdatedAt:   dataBooking.UpdatedAt,
		UserId:      user.Id,
		User:        *user,
		PlaceId:     dataBooking.PlaceId,
		Place:       *place,
		StatusId:    dataBooking.StatusId,
		CheckInDate: dataBooking.CheckInDate,
		ChekoutDate: dataBooking.ChekoutDate,
	}
}
