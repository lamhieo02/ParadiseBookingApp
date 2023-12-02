package bookinghandler

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
)

type BookingUseCase interface {
	ListBooking(ctx context.Context, paging *common.Paging, filter *iomodel.FilterListBooking, userID int) (result []entities.Booking, err error)
	CreateBooking(ctx context.Context, bookingData *iomodel.CreateBookingReq) error
	UpdateStatusBooking(ctx context.Context, bookingID, status int) error
}

type bookingHandler struct {
	bookingUC BookingUseCase
}

func NewBookingHandler(bookingUC BookingUseCase) *bookingHandler {
	return &bookingHandler{bookingUC: bookingUC}
}
