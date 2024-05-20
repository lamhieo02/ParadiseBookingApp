package bookingratingusecase

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"paradise-booking/modules/booking_rating/iomodel"
)

func (uc *bookingRatingUsecase) GetCommentByVendorID(ctx context.Context, vendorID int, objectType int) (*iomodel.GetCommentByVendorResp, error) {
	res, err := uc.BookingRatingSto.GetByVendorID(ctx, vendorID, objectType)
	if err != nil {
		return nil, err
	}

	var result iomodel.GetCommentByVendorResp
	for _, bookingRate := range res {

		user, err := uc.AccountSto.GetProfileByID(ctx, bookingRate.UserId)
		if err != nil {
			log.Printf("Error when get user profile by id: %v\n", err)
			return nil, common.ErrCannotGetEntity(entities.Account{}.TableName(), err)
		}

		if objectType == constant.BookingRatingObjectTypePlace {
			place, err := uc.PlaceSto.GetPlaceByID(ctx, bookingRate.ObjectId)
			if err != nil {
				log.Printf("Error when get place by id: %v\n", err)
				continue
			}
			result.ListRating = append(result.ListRating, iomodel.GetCommentUserByVendor{
				DataRating: bookingRate,
				DataPlace:  place,
				DataUser:   *user,
			})
		} else if objectType == constant.BookingRatingObjectTypeGuide {
			postGuide, err := uc.PostGuideSto.GetByID(ctx, bookingRate.ObjectId)
			if err != nil {
				log.Printf("Error when get post guide by id: %v\n", err)
				continue
			}
			result.ListRating = append(result.ListRating, iomodel.GetCommentUserByVendor{
				DataRating:    bookingRate,
				DataPostGuide: postGuide,
				DataUser:      *user,
			})
		}
	}

	// dataVendor, err := uc.AccountSto.GetProfileByID(ctx, vendorID)
	// if err != nil {
	// 	return nil, common.ErrCannotGetEntity(entities.Account{}.TableName(), err)
	// }

	// result.DataVendor = dataVendor

	return &result, nil
}
