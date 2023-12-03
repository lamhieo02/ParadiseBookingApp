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
	GetByID(ctx context.Context, id int) (*entities.Booking, error)
	GetByPlaceID(ctx context.Context, placeId int) ([]entities.Booking, error)
	ListPlaceIds(ctx context.Context) ([]int, error)
}

type BookingDetailStorage interface {
	Create(ctx context.Context, data *entities.BookingDetail) (err error)
	CreateTx(ctx context.Context, createBookingDetailTxParam bookingdetailstorage.CreateBookingDetailTxParam) error
}

type AccountSto interface {
	GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error)
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type PlaceSto interface {
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
	ListPlaceNotInIds(ctx context.Context, placeIds []int, vendorId int) ([]entities.Place, error)
}

type bookingUseCase struct {
	bookingSto       BookingStorage
	bookingDetailSto BookingDetailStorage
	AccountSto       AccountSto
	cfg              *config.Config
	taskDistributor  worker.TaskDistributor
	PlaceSto         PlaceSto
}

func NewBookingUseCase(bookingStore BookingStorage, bookingDetailStorage BookingDetailStorage, config *config.Config, taskDistributor worker.TaskDistributor, accountSto AccountSto, placeSto PlaceSto) *bookingUseCase {
	return &bookingUseCase{bookingSto: bookingStore, bookingDetailSto: bookingDetailStorage, cfg: config, taskDistributor: taskDistributor, AccountSto: accountSto, PlaceSto: placeSto}
}
