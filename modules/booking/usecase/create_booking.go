package bookingusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"paradise-booking/modules/booking/convert"
	"paradise-booking/modules/booking/iomodel"
	bookingdetailstorage "paradise-booking/modules/booking_detail/storage"
	"paradise-booking/worker"
	"time"

	"github.com/hibiken/asynq"
)

func (uc *bookingUseCase) CreateBooking(ctx context.Context, bookingData *iomodel.CreateBookingReq) error {
	// convert iomodel to entities
	bookingEntity := convert.ConvertBookingModelToBookingEntity(bookingData)
	bookingDetailEntity := convert.ConvertBookingModelToBookingDetail(bookingData)

	bookingEntity.StatusId = constant.BookingStatusPending // default status is pending when create booking
	// create booking
	if err := uc.bookingSto.Create(ctx, bookingEntity); err != nil {
		return common.ErrCannotCreateEntity(bookingEntity.TableName(), err)
	}

	bookingDetailEntity.BookingId = bookingEntity.Id

	// create booking detail and send mail to customer to confirm booking in 1 transaction
	paramCreateTx := bookingdetailstorage.CreateBookingDetailTxParam{
		Data: bookingDetailEntity,
		AfterCreate: func(data *entities.BookingDetail) error {
			// after create booking success, we will send email to user to verify booking
			taskPayload := worker.PayloadSendConfirmBooking{
				BookingID: bookingEntity.Id,
				Email:     bookingDetailEntity.Email,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueSendVerifyEmail),
			}

			return uc.taskDistributor.DistributeTaskSendConfirmBooking(ctx, &taskPayload, opts...)
		},
	}

	if err := uc.bookingDetailSto.CreateTx(ctx, paramCreateTx); err != nil {
		return err
	}

	return nil
}
