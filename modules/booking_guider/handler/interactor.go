package bookingguiderhandler

import (
	"context"
	"paradise-booking/common"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
)

type bookingGuiderUseCase interface {
	CreateBookingGuider(ctx context.Context, bookingData *bookingguideriomodel.CreateBookingReq) (*bookingguideriomodel.CreateBookingResp, error)
	UpdateStatusBooking(ctx context.Context, bookingGuiderID, status int) error
	GetBookingByID(ctx context.Context, bookingGuiderID int) (*bookingguideriomodel.GetBookingGuiderResp, error)
	GetBookingByUserID(ctx context.Context, userID int) ([]*bookingguideriomodel.GetBookingGuiderResp, error)
	ListBooking(ctx context.Context, paging *common.Paging, filter *bookingguideriomodel.Filter) ([]*bookingguideriomodel.GetBookingGuiderResp, error)
	DeleteBookingByID(ctx context.Context, bookingGuiderID int) error
}

type bookingGuiderHandler struct {
	bookingGuiderUC bookingGuiderUseCase
}

func NewBookingGuiderHandler(bookingGuiderUseCase bookingGuiderUseCase) *bookingGuiderHandler {
	return &bookingGuiderHandler{bookingGuiderUseCase}
}
