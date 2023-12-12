package bookingratingusecase

import (
	"context"
	"paradise-booking/entities"
)

type BookingRatingSto interface {
	Create(ctx context.Context, data *entities.BookingRating) (*entities.BookingRating, error)
	GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.BookingRating, error)
	GetByVendorID(ctx context.Context, vendorID int) ([]entities.BookingRating, error)
}

type AccountSto interface {
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type PlaceSto interface {
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
}

type bookingRatingUsecase struct {
	BookingRatingSto BookingRatingSto
	AccountSto       AccountSto
	PlaceSto         PlaceSto
}

func Newbookingratingusecase(BookingRatingSto BookingRatingSto, accountSto AccountSto, placeSto PlaceSto) *bookingRatingUsecase {
	return &bookingRatingUsecase{BookingRatingSto, accountSto, placeSto}
}
