package bookingguiderusecase

import (
	"context"
	"paradise-booking/constant"
	"paradise-booking/entities"
	bookingguiderconvert "paradise-booking/modules/booking_guider/convert"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
	momoprovider "paradise-booking/provider/momo"
	"paradise-booking/worker"
	"time"

	"github.com/hibiken/asynq"
)

const (
	bookingGuider = "Booking Guider"
)

func (uc *bookingGuiderUseCase) CreateBookingGuider(ctx context.Context, bookingData *bookingguideriomodel.CreateBookingReq) (*bookingguideriomodel.CreateBookingResp, error) {
	var err error
	bookingEntity := bookingData.ToEntity()

	bookingEntity.StatusID = constant.BookingGuiderStatusPending
	// create booking
	if err := uc.bookingGuiderSto.Create(ctx, bookingEntity); err != nil {
		return nil, err
	}

	if err := uc.sendMailToConfirmBookingGuider(ctx, bookingEntity); err != nil {
		return nil, err
	}

	// if payment method is momo, we will create payment with momo
	var requestId, orderId, paymentUrl string
	if bookingData.PaymentMethod == constant.PaymentMethodMomo {
		bookingInfo := &momoprovider.InfoPayment{
			BookingID:   bookingEntity.Id,
			NameBooking: bookingGuider,
			Email:       bookingEntity.Email,
			TotalPrice:  bookingEntity.TotalPrice,
			RedirectURL: constant.RedirectURLBookingGuiderMomo,
		}

		orderId, requestId, paymentUrl, err = uc.momoProvider.CreatePayment(bookingInfo)
		if err != nil {
			return nil, err
		}
	}

	// create data payment
	status := constant.PaymentStatusUnpaid
	if bookingData.PaymentMethod == constant.PaymentMethodMomo {
		status = constant.PaymentStatusPaid
	}

	paymentEntity := &entities.Payment{
		BookingID: bookingEntity.Id,
		MethodID:  bookingData.PaymentMethod,
		StatusID:  status,
		Amount:    bookingEntity.TotalPrice,
		RequestID: requestId,
		OrderID:   orderId,
		Type:      constant.PaymentTypeBookingGuider,
	}

	err = uc.paymentSto.CreatePayment(ctx, paymentEntity)
	if err != nil {
		return nil, err
	}

	// get booking guider data
	result := bookingguideriomodel.CreateBookingResp{
		PaymentUrl:        paymentUrl,
		BookingGuiderData: *bookingguiderconvert.ConvertBookingEntityToModel(bookingEntity),
	}

	return &result, nil
}

func (uc *bookingGuiderUseCase) sendMailToConfirmBookingGuider(ctx context.Context, bookingEntity *entities.BookingGuider) error {
	// send mail to customer to confirm booking guider
	taskPayload := worker.PayloadSendConfirmBookingGuider{
		BookingGuiderID: bookingEntity.Id,
		Email:           bookingEntity.Email,
		FullName:        bookingEntity.Name,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueSendVerifyEmail),
	}

	if err := uc.taskDistributor.DistributeTaskSendConfirmBookingGuider(ctx, &taskPayload, opts...); err != nil {
		return err
	}

	return nil
}
