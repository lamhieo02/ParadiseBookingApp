package bookingguiderconvert

import (
	"paradise-booking/constant"
	"paradise-booking/entities"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
)

func ConvertBookingEntityToModel(entity *entities.BookingGuider, postGuide *postguideiomodel.GetPostGuideResp) *bookingguideriomodel.GetBookingGuiderResp {
	res := &bookingguideriomodel.GetBookingGuiderResp{
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
		PaymentMethod:    constant.MapPaymentMethod[entity.PaymentMethod],
	}

	if entity.CreatedAt != nil {
		res.CreatedAt = *entity.CreatedAt
	}

	if postGuide != nil {
		res.PostGuide = *postGuide
	}
	return res
}
