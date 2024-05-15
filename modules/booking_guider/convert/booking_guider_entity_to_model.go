package bookingguiderconvert

import (
	"paradise-booking/constant"
	"paradise-booking/entities"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
)

func ConvertBookingEntityToModel(entity *entities.BookingGuider) *bookingguideriomodel.GetBookingGuiderResp {
	return &bookingguideriomodel.GetBookingGuiderResp{
		ID:               entity.Id,
		CalendarGuiderID: entity.CalendarGuiderID,
		Email:            entity.Email,
		NumberOfPeople:   entity.NumberOfPeople,
		Name:             entity.Name,
		Note:             entity.Note,
		StatusID:         entity.StatusID,
		Status:           constant.MapBookingGuiderStatus[entity.StatusID],
		TotalPrice:       int(entity.TotalPrice),
		Phone:            entity.Phone,
		CreatedAt:        *entity.CreatedAt,
		PaymentMethod:    constant.MapPaymentMethod[entity.PaymentMethod],
	}
}
