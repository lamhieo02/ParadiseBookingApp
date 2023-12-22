package paymentusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (uc *paymentUseCase) ListPaymentByVendorID(ctx context.Context, paging *common.Paging, vendorID int) ([]entities.Payment, error) {

	payments, err := uc.paymentSto.GetPaymentByVendor(ctx, int(vendorID), paging)
	if err != nil {
		return nil, err
	}

	return payments, nil
}
