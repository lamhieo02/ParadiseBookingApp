package bookingguiderusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (uc *bookingGuiderUseCase) UpdateStatusBooking(ctx context.Context, bookingGuiderID, status int) error {
	// update status booking
	if err := uc.bookingGuiderSto.UpdateStatus(ctx, bookingGuiderID, status); err != nil {
		return common.ErrCannotUpdateEntity(entities.Booking{}.TableName(), err)
	}

	return nil
}
