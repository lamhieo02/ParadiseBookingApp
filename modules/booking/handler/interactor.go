package bookinghandler

import (
	"context"
	"paradise-booking/modules/booking/iomodel"
)

type BookingUseCase interface {
	CreateBooking(ctx context.Context, bookingData *iomodel.CreateBookingReq) error
	UpdateStatusBooking(ctx context.Context, bookingID, status int) error
}

type bookingHandler struct {
	bookingUC BookingUseCase
}

func NewBookingHandler(bookingUC BookingUseCase) *bookingHandler {
	return &bookingHandler{bookingUC: bookingUC}
}
