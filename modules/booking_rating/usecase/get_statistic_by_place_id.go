package bookingratingusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *bookingRatingUsecase) GetStatisticByObjectID(ctx context.Context, objectID int, objectType int) ([]entities.StatisticResp, error) {
	res, err := uc.BookingRatingSto.GetStatisticByObjectID(ctx, int64(objectID), objectType)
	if err != nil {
		return nil, err
	}

	return res, nil
}
