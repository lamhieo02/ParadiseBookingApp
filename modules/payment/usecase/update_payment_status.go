package paymentusecase

import (
	"context"
)

func (uc *paymentUseCase) UpdateStatusPaymentByID(ctx context.Context, id int, statusID int) error {
	payment, err := uc.paymentSto.GetByID(ctx, id)
	if err != nil {
		return err
	}

	payment.StatusID = statusID

	if err := uc.paymentSto.UpdateByID(ctx, id, payment); err != nil {
		return err
	}

	return nil
}
