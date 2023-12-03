package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
)

func ConvertBookingModelToGetResp(user *entities.Account, dataBooking *entities.Booking, place *entities.Place) *iomodel.GetBookingResp {
	return &iomodel.GetBookingResp{
		UserId: user.Id,
		User:   *user,
		GetData: iomodel.DataListBooking{
			Id:          dataBooking.Id,
			CreatedAt:   dataBooking.CreatedAt,
			UpdatedAt:   dataBooking.UpdatedAt,
			PlaceId:     dataBooking.PlaceId,
			Place:       *place,
			StatusId:    dataBooking.StatusId,
			CheckInDate: dataBooking.CheckInDate,
			ChekoutDate: dataBooking.ChekoutDate,
		},
	}
}
