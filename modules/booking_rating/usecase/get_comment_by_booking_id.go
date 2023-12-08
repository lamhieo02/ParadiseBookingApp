package bookingratingusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (uc *bookingRatingUsecase) GetCommentByBookingID(ctx context.Context, bookingID int) ([]entities.BookingRating, error) {
	res, err := uc.BookingRatingSto.GetByCondition(ctx, map[string]interface{}{"booking_id": bookingID})
	if err != nil {
		return nil, common.ErrCannotGetEntity(entities.BookingRating{}.TableName(), err)
	}

	return res, nil
}
