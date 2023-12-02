package bookingusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
)

func (uc *bookingUseCase) ListBooking(ctx context.Context, paging *common.Paging, filter *iomodel.FilterListBooking, userID int) (result []entities.Booking, err error) {

	// Assign permissions by userid
	if paging != nil {
		paging.Process()
	}
	result, err = uc.bookingSto.ListByFilter(ctx, filter, paging, userID)

	if err != nil {
		return nil, common.ErrCannotListEntity("booking", err)
	}

	return result, nil
}
