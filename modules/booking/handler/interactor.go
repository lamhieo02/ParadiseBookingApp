package bookinghandler

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
)

type BookingUseCase interface {
	ListBooking(ctx context.Context, paging *common.Paging, filter *iomodel.FilterListBooking, userID int) (*iomodel.ListBookingResp, error)
	CreateBooking(ctx context.Context, bookingData *iomodel.CreateBookingReq) (*entities.Booking, error)
	UpdateStatusBooking(ctx context.Context, bookingID, status int) error
	GetBookingByID(ctx context.Context, id int) (*iomodel.GetBookingResp, error)
	GetBookingByPlaceID(ctx context.Context, placeId int) ([]iomodel.GetBookingByPlaceResp, error)
	ListPlaceReservationByVendor(ctx context.Context, vendorId, typeManage int) (res []entities.Place, err error)
	ListBookingByCondition(ctx context.Context) ([]entities.Booking, error)
	DeleteBookingByID(ctx context.Context, id int) error
}

type bookingHandler struct {
	bookingUC BookingUseCase
}

func NewBookingHandler(bookingUC BookingUseCase) *bookingHandler {
	return &bookingHandler{bookingUC: bookingUC}
}
