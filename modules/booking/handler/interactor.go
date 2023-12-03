package bookinghandler

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/booking/iomodel"
)

type BookingUseCase interface {
	ListBooking(ctx context.Context, paging *common.Paging, filter *iomodel.FilterListBooking, userID int) (*iomodel.ListBookingResp, error)
	CreateBooking(ctx context.Context, bookingData *iomodel.CreateBookingReq) error
	UpdateStatusBooking(ctx context.Context, bookingID, status int) error
}

type bookingHandler struct {
	bookingUC BookingUseCase
}

func NewBookingHandler(bookingUC BookingUseCase) *bookingHandler {
	return &bookingHandler{bookingUC: bookingUC}
}
