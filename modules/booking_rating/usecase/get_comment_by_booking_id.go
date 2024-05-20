package bookingratingusecase

import (
	"context"
	"log"
	"paradise-booking/constant"
	"paradise-booking/modules/booking_rating/iomodel"
)

func (uc *bookingRatingUsecase) GetCommentByBookingID(ctx context.Context, bookingID int, objectType int) ([]iomodel.GetCommentResp, error) {
	res, err := uc.BookingRatingSto.GetByCondition(ctx, map[string]interface{}{"booking_id": bookingID, "object_type": objectType})
	if err != nil {
		return nil, err
	}

	var result []iomodel.GetCommentResp

	for _, bookingRate := range res {
		user, err := uc.AccountSto.GetProfileByID(ctx, bookingRate.UserId)
		if err != nil {
			log.Printf("Error when get user profile by id: %v\n", err)
			continue
		}

		if objectType == constant.BookingRatingObjectTypePlace {
			place, err := uc.PlaceSto.GetPlaceByID(ctx, bookingRate.ObjectId)
			if err != nil {
				log.Printf("Error when get place by id: %v\n", err)
				continue
			}
			result = append(result, iomodel.GetCommentResp{
				DataRating: bookingRate,
				DataUser:   *user,
				DataPlace:  place,
			})
		} else if objectType == constant.BookingRatingObjectTypeGuide {
			postGuide, err := uc.PostGuideSto.GetByID(ctx, bookingRate.ObjectId)
			if err != nil {
				log.Printf("Error when get post guide by id: %v\n", err)
				continue
			}

			result = append(result, iomodel.GetCommentResp{
				DataRating:    bookingRate,
				DataUser:      *user,
				DataPostGuide: postGuide,
			})
		}
	}

	return result, nil
}
