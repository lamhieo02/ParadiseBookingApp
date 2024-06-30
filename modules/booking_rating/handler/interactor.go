package bookingratinghandler

import (
	"context"
	"paradise-booking/entities"
	bookingratingiomodel "paradise-booking/modules/booking_rating/iomodel"
)

type placeRatingUseCase interface {
	MakeComment(ctx context.Context, userID int, data *bookingratingiomodel.CreateBookingRatingReq) (*entities.BookingRating, error)
	GetCommentByObjectID(ctx context.Context, objectID int, objectType int) (*bookingratingiomodel.GetCommentByObjectResp, error)
	GetCommentByBookingID(ctx context.Context, bookingID int, objectType int) ([]bookingratingiomodel.GetCommentResp, error)
	GetCommentByUserID(ctx context.Context, usrID int, objectType int) (*bookingratingiomodel.GetCommentByUserResp, error)
	GetCommentByVendorID(ctx context.Context, vendorID int, objectType int) (*bookingratingiomodel.GetCommentByVendorResp, error)
	GetStatisticByObjectID(ctx context.Context, objectID int, objectType int) ([]entities.StatisticResp, error)
}

type bookingratinghandler struct {
	placeRatingUC placeRatingUseCase
}

func Newbookingratinghandler(placeRatingUseCase placeRatingUseCase) *bookingratinghandler {
	return &bookingratinghandler{placeRatingUC: placeRatingUseCase}
}
