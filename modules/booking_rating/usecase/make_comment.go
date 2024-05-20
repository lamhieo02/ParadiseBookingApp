package bookingratingusecase

import (
	"context"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"paradise-booking/modules/booking_rating/iomodel"
)

func (u *bookingRatingUsecase) MakeComment(ctx context.Context, userID int, data *iomodel.CreateBookingRatingReq) (*entities.BookingRating, error) {

	model := entities.BookingRating{
		UserId:    userID,
		BookingId: data.BookingID,
		Title:     data.Title,
		Content:   data.Content,
		Rating:    int(data.Rating),
		// PlaceId:   data.PlaceID,
		ObjectId:   data.ObjectID,
		ObjectType: data.ObjectType,
	}

	if _, err := u.BookingRatingSto.Create(ctx, &model); err != nil {
		return nil, err
	}

	if data.ObjectType == constant.BookingRatingObjectTypePlace {
		// delete rating of object in cache
		place := entities.Place{}
		place.Id = int(data.ObjectID)
		key := place.CacheKeyPlaceRating()
		u.cache.Delete(ctx, key)
	} else if data.ObjectType == constant.BookingRatingObjectTypeGuide {
		// delete rating of object in cache
		guide := entities.PostGuide{}
		guide.Id = int(data.ObjectID)
		key := guide.CacheKeyGuideRating()
		u.cache.Delete(ctx, key)
	}

	return &model, nil
}
