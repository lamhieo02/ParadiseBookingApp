package bookingratingusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (uc *bookingRatingUsecase) GetCommentByPlaceID(ctx context.Context, placeID int) ([]entities.BookingRating, error) {
	res, err := uc.BookingRatingSto.GetByCondition(ctx, map[string]interface{}{"place_id": placeID})
	if err != nil {
		return nil, common.ErrCannotGetEntity(entities.BookingRating{}.TableName(), err)
	}

	return res, nil
}
