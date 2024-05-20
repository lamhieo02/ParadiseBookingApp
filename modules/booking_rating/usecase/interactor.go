package bookingratingusecase

import (
	"context"
	"paradise-booking/entities"
)

type BookingRatingSto interface {
	Create(ctx context.Context, data *entities.BookingRating) (*entities.BookingRating, error)
	GetByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.BookingRating, error)
	GetByVendorID(ctx context.Context, vendorID int, objectType int) ([]entities.BookingRating, error)
	GetStatisticByObjectID(ctx context.Context, objectId int64, objectType int) ([]entities.StatisticResp, error)
}

type AccountSto interface {
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type PlaceSto interface {
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
}

type Cache interface {
	Delete(ctx context.Context, key string)
}

type PostGuideSto interface {
	GetByID(ctx context.Context, id int) (*entities.PostGuide, error)
}

type bookingRatingUsecase struct {
	BookingRatingSto BookingRatingSto
	AccountSto       AccountSto
	PlaceSto         PlaceSto
	cache            Cache
	PostGuideSto     PostGuideSto
}

func Newbookingratingusecase(BookingRatingSto BookingRatingSto, accountSto AccountSto, placeSto PlaceSto, cache Cache, postGuideSto PostGuideSto) *bookingRatingUsecase {
	return &bookingRatingUsecase{BookingRatingSto, accountSto, placeSto, cache, postGuideSto}
}
