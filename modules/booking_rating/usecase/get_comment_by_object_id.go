package bookingratingusecase

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"paradise-booking/modules/booking_rating/iomodel"
)

func (uc *bookingRatingUsecase) GetCommentByObjectID(ctx context.Context, objectID int, objectType int) (*iomodel.GetCommentByObjectResp, error) {
	res, err := uc.BookingRatingSto.GetByCondition(ctx, map[string]interface{}{"object_id": objectID, "object_type": objectType})
	if err != nil {
		return nil, err
	}

	var result iomodel.GetCommentByObjectResp
	var listRating []iomodel.GetCommentRespByObject
	for _, bookingRate := range res {
		user, err := uc.AccountSto.GetProfileByID(ctx, bookingRate.UserId)
		if err != nil {
			log.Printf("Error when get user profile by id: %v\n", err)
			continue
		}

		listRating = append(listRating, iomodel.GetCommentRespByObject{
			DataRating: bookingRate,
			DataUser:   *user,
		})
	}
	result.ListRating = listRating

	if objectType == constant.BookingRatingObjectTypePlace {
		place, err := uc.PlaceSto.GetPlaceByID(ctx, objectID)
		if err != nil {
			log.Printf("Error when get place by id: %v\n", err)
			return nil, common.ErrCannotGetEntity(entities.Place{}.TableName(), err)
		}
		result.DataPlace = place
	} else if objectType == constant.BookingRatingObjectTypeGuide {
		postGuide, err := uc.PostGuideSto.GetByID(ctx, objectID)
		if err != nil {
			log.Printf("Error when get post guide by id: %v\n", err)
			return nil, common.ErrCannotGetEntity(entities.PostGuide{}.TableName(), err)
		}
		result.DataPostGuide = postGuide
	}

	return &result, nil
}
