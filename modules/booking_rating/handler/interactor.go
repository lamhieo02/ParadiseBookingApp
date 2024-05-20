package bookingratinghandler

import (
	"context"
	"paradise-booking/entities"
	"paradise-booking/modules/booking_rating/iomodel"
)

type placeRatingUseCase interface {
	MakeComment(ctx context.Context, userID int, data *iomodel.CreateBookingRatingReq) (*entities.BookingRating, error)
	GetCommentByObjectID(ctx context.Context, objectID int, objectType int) (*iomodel.GetCommentByObjectResp, error)
	GetCommentByBookingID(ctx context.Context, bookingID int, objectType int) ([]iomodel.GetCommentResp, error)
	GetCommentByUserID(ctx context.Context, usrID int, objectType int) (*iomodel.GetCommentByUserResp, error)
	GetCommentByVendorID(ctx context.Context, vendorID int, objectType int) (*iomodel.GetCommentByVendorResp, error)
	GetStatisticByObjectID(ctx context.Context, objectID int, objectType int) ([]entities.StatisticResp, error)
}

type bookingratinghandler struct {
	placeRatingUC placeRatingUseCase
}

func Newbookingratinghandler(placeRatingUseCase placeRatingUseCase) *bookingratinghandler {
	return &bookingratinghandler{placeRatingUC: placeRatingUseCase}
}
