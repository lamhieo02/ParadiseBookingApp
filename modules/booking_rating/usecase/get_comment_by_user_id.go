package bookingratingusecase

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	bookingratingconvert "paradise-booking/modules/booking_rating/convert"
	bookingratingiomodel "paradise-booking/modules/booking_rating/iomodel"
)

func (uc *bookingRatingUsecase) GetCommentByUserID(ctx context.Context, usrID int, objectType int) (*bookingratingiomodel.GetCommentByUserResp, error) {
	res, err := uc.BookingRatingSto.GetByCondition(ctx, map[string]interface{}{"user_id": usrID, "object_type": objectType})
	if err != nil {
		return nil, common.ErrCannotGetEntity(entities.BookingRating{}.TableName(), err)
	}

	var result bookingratingiomodel.GetCommentByUserResp
	for _, bookingRate := range res {

		if objectType == constant.BookingRatingObjectTypePlace {
			place, err := uc.PlaceSto.GetPlaceByID(ctx, bookingRate.ObjectId)
			if err != nil {
				log.Printf("Error when get place by id: %v\n", err)
				continue
			}
			result.ListRating = append(result.ListRating, bookingratingiomodel.GetCommentRespByUser{
				DataRating: bookingratingconvert.ConvertDataBookingRatingEntityToModel(&bookingRate),
				DataPlace:  bookingratingconvert.ConvertPlaceEntityToModel(place),
			})
		} else if objectType == constant.BookingRatingObjectTypeGuide {
			postGuide, err := uc.PostGuideSto.GetByID(ctx, bookingRate.ObjectId)
			if err != nil {
				log.Printf("Error when get post guide by id: %v\n", err)
				continue
			}
			result.ListRating = append(result.ListRating, bookingratingiomodel.GetCommentRespByUser{
				DataRating:    bookingratingconvert.ConvertDataBookingRatingEntityToModel(&bookingRate),
				DataPostGuide: bookingratingconvert.ConvertPostGuideEntityToModel(postGuide),
			})
		}
	}

	user, err := uc.AccountSto.GetProfileByID(ctx, usrID)
	if err != nil {
		log.Printf("Error when get user profile by id: %v\n", err)
		return nil, common.ErrCannotGetEntity(entities.Account{}.TableName(), err)
	}
	result.DataUser = *user

	return &result, nil
}
