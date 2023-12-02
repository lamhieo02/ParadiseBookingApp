package bookingusecase

import (
	"context"
	"paradise-booking/config"
	"paradise-booking/entities"
	bookingdetailstorage "paradise-booking/modules/booking_detail/storage"
	"paradise-booking/worker"
)

type BookingStorage interface {
	Create(ctx context.Context, data *entities.Booking) (err error)
	UpdateStatus(ctx context.Context, bookingID int, status int) error
}

type BookingDetailStorage interface {
	Create(ctx context.Context, data *entities.BookingDetail) (err error)
	CreateTx(ctx context.Context, createBookingDetailTxParam bookingdetailstorage.CreateBookingDetailTxParam) error
}

type bookingUseCase struct {
	bookingSto       BookingStorage
	bookingDetailSto BookingDetailStorage
	cfg              *config.Config
	taskDistributor  worker.TaskDistributor
}

func NewBookingUseCase(bookingStore BookingStorage, bookingDetailStorage BookingDetailStorage, config *config.Config, taskDistributor worker.TaskDistributor) *bookingUseCase {
	return &bookingUseCase{bookingSto: bookingStore, bookingDetailSto: bookingDetailStorage, cfg: config, taskDistributor: taskDistributor}
}
