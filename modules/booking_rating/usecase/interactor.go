package bookingratingusecase

import (
	"context"
	"paradise-booking/entities"
)

type BookingRatingSto interface {
	Create(ctx context.Context, data *entities.BookingRating) (*entities.BookingRating, error)
	GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.BookingRating, error)
}

type AccountSto interface {
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type bookingRatingUsecase struct {
	BookingRatingSto BookingRatingSto
	AccountSto       AccountSto
}

func Newbookingratingusecase(BookingRatingSto BookingRatingSto, accountSto AccountSto) *bookingRatingUsecase {
	return &bookingRatingUsecase{BookingRatingSto, accountSto}
}
