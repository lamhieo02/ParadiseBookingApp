package bookingguiderusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
	momoprovider "paradise-booking/provider/momo"
	"paradise-booking/worker"
)

type bookingGuiderStorage interface {
	Create(ctx context.Context, data *entities.BookingGuider) error
	UpdateWithMap(ctx context.Context, id int, props map[string]interface{}) error
	UpdateStatus(ctx context.Context, bookingGuiderID int, status int) error
	GetByID(ctx context.Context, id int) (*entities.BookingGuider, error)
	ListByCondition(ctx context.Context, conditions []common.Condition) ([]*entities.BookingGuider, error)
	ListByFilter(ctx context.Context, paging *common.Paging, filter *bookingguideriomodel.Filter, userId int) ([]entities.BookingGuider, error)
}

type PaymentSto interface {
	CreatePayment(ctx context.Context, payment *entities.Payment) error
}

type CalendarSto interface {
	GetByID(ctx context.Context, id int) (*entities.CalendarGuider, error)
}

type bookingGuiderUseCase struct {
	bookingGuiderSto bookingGuiderStorage
	taskDistributor  worker.TaskDistributor
	momoProvider     *momoprovider.Momo
	paymentSto       PaymentSto
	calendarSto      CalendarSto
}

func NewBookingGuiderUseCase(bookingGuiderSto bookingGuiderStorage, taskDistributor worker.TaskDistributor, momoProvider *momoprovider.Momo, paymentSto PaymentSto, calendarSto CalendarSto) *bookingGuiderUseCase {
	return &bookingGuiderUseCase{bookingGuiderSto, taskDistributor, momoProvider, paymentSto, calendarSto}
}
