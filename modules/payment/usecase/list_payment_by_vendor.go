package paymentusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"

	"github.com/samber/lo"
)

func (uc *paymentUseCase) ListPaymentByVendorID(ctx context.Context, paging *common.Paging, vendorID int, bookingId int) ([]entities.Payment, error) {

	paging.Process()
	payments, err := uc.paymentSto.GetPaymentByVendor(ctx, int(vendorID), paging)
	if err != nil {
		return nil, err
	}

	if bookingId != 0 {
		res := lo.Filter(payments, func(item entities.Payment, _ int) bool {
			return item.BookingID == bookingId
		})
		return res, nil
	}

	return payments, nil
}