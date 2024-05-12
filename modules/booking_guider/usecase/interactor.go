package bookingguiderusecase

import (
	"context"
	"paradise-booking/entities"
	momoprovider "paradise-booking/provider/momo"
	"paradise-booking/worker"
)

type bookingGuiderStorage interface {
	Create(ctx context.Context, data *entities.BookingGuider) error
	UpdateWithMap(ctx context.Context, id int, props map[string]interface{}) error
	UpdateStatus(ctx context.Context, bookingGuiderID int, status int) error
}

type PaymentSto interface {
	CreatePayment(ctx context.Context, payment *entities.Payment) error
}

type bookingGuiderUseCase struct {
	bookingGuiderSto bookingGuiderStorage
	taskDistributor  worker.TaskDistributor
	momoProvider     *momoprovider.Momo
	paymentSto       PaymentSto
}

func NewBookingGuiderUseCase(bookingGuiderSto bookingGuiderStorage, taskDistributor worker.TaskDistributor, momoProvider *momoprovider.Momo, paymentSto PaymentSto) *bookingGuiderUseCase {
	return &bookingGuiderUseCase{bookingGuiderSto, taskDistributor, momoProvider, paymentSto}
}
