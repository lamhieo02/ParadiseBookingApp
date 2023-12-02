package bookingusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/config"
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
	bookingdetailstorage "paradise-booking/modules/booking_detail/storage"
	"paradise-booking/worker"
)

type BookingStorage interface {
	Create(ctx context.Context, data *entities.Booking) (err error)
	UpdateStatus(ctx context.Context, bookingID int, status int) error
	ListByFilter(ctx context.Context, filter *iomodel.FilterListBooking, paging *common.Paging, userId int) ([]entities.Booking, error)
}

type BookingDetailStorage interface {
	Create(ctx context.Context, data *entities.BookingDetail) (err error)
	CreateTx(ctx context.Context, createBookingDetailTxParam bookingdetailstorage.CreateBookingDetailTxParam) error
}

type AccountSto interface {
	GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error)
}

type bookingUseCase struct {
	bookingSto       BookingStorage
	bookingDetailSto BookingDetailStorage
	AccountSto       AccountSto
	cfg              *config.Config
	taskDistributor  worker.TaskDistributor
}

func NewBookingUseCase(bookingStore BookingStorage, bookingDetailStorage BookingDetailStorage, config *config.Config, taskDistributor worker.TaskDistributor, accountSto AccountSto) *bookingUseCase {
	return &bookingUseCase{bookingSto: bookingStore, bookingDetailSto: bookingDetailStorage, cfg: config, taskDistributor: taskDistributor, AccountSto: accountSto}
}
