package bookingguiderusecase

import (
	"context"
)

func (uc *bookingGuiderUseCase) DeleteBookingByID(ctx context.Context, bookingGuiderID int) error {

	err := uc.bookingGuiderSto.DeleteByID(ctx, bookingGuiderID)
	if err != nil {
		return err
	}

	return nil

}
