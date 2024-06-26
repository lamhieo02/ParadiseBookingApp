package paymentusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"

	"github.com/samber/lo"
)

func (uc *paymentUseCase) ListPaymentByGuiderID(ctx context.Context, paging *common.Paging, guiderID int, bookingId int) ([]entities.Payment, error) {

	if paging.Limit == 0 {
		paging.Limit = constant.PaymentPagingLimitMax
		paging.Page = constant.PaymentPagingPageDefault
	}

	payments, err := uc.paymentSto.GetPaymentByGuider(ctx, int(guiderID), paging)
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
