package bookingguiderhandler

import (
	"context"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
)

type bookingGuiderUseCase interface {
	CreateBookingGuider(ctx context.Context, bookingData *bookingguideriomodel.CreateBookingReq) (*bookingguideriomodel.CreateBookingResp, error)
	UpdateStatusBooking(ctx context.Context, bookingGuiderID, status int) error
	GetBookingByID(ctx context.Context, bookingGuiderID int) (*bookingguideriomodel.GetBookingGuiderResp, error)
}

type bookingGuiderHandler struct {
	bookingGuiderUC bookingGuiderUseCase
}

func NewBookingGuiderHandler(bookingGuiderUseCase bookingGuiderUseCase) *bookingGuiderHandler {
	return &bookingGuiderHandler{bookingGuiderUseCase}
}
