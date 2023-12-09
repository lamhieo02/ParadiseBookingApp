package bookingratingusecase

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/booking_rating/iomodel"
)

func (uc *bookingRatingUsecase) GetCommentByUserID(ctx context.Context, usrID int) (*iomodel.GetCommentByUserResp, error) {
	res, err := uc.BookingRatingSto.GetByCondition(ctx, map[string]interface{}{"user_id": usrID})
	if err != nil {
		return nil, common.ErrCannotGetEntity(entities.BookingRating{}.TableName(), err)
	}

	var result iomodel.GetCommentByUserResp

	result.DataRating = append(result.DataRating, res...)
	user, err := uc.AccountSto.GetProfileByID(ctx, usrID)
	if err != nil {
		log.Printf("Error when get user profile by id: %v\n", err)
		return nil, common.ErrCannotGetEntity(entities.Account{}.TableName(), err)
	}
	result.DataUser = *user

	return &result, nil
}
