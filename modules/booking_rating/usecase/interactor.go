package bookingratingusecase

import (
	"context"
	"paradise-booking/entities"
)

type BookingRatingSto interface {
	Create(ctx context.Context, data *entities.BookingRating) (*entities.BookingRating, error)
	GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.BookingRating, error)
}

type bookingRatingUsecase struct {
	BookingRatingSto BookingRatingSto
}

func Newbookingratingusecase(BookingRatingSto BookingRatingSto) *bookingRatingUsecase {
	return &bookingRatingUsecase{BookingRatingSto}
}
