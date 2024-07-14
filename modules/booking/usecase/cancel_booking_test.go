package bookingusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCancelBooking(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockBookingStorage := NewMockBookingStorage(ctrl)

	// init usecase
	cfg := config.Config{}
	uc := NewBookingUseCase(mockBookingStorage, nil, &cfg, nil, nil, nil, nil, nil)

	bookingId := gofakeit.Number(1, 100)

	Convey("Test Cancel Booking", t, func() {
		Convey("Get Booking by id fail", func() {
			mockBookingStorage.EXPECT().GetByID(ctx, bookingId).Return(nil, errors.New("error"))
			err := uc.CancelBooking(ctx, bookingId)
			So(err, ShouldNotBeNil)
		})
		Convey("Get Booking by id success", func() {
			mockBookingStorage.EXPECT().GetByID(ctx, bookingId).Return(&entities.Booking{StatusId: constant.BookingStatusPending}, nil)
			Convey("Update Booking fail", func() {
				mockBookingStorage.EXPECT().UpdateStatus(ctx, bookingId, gomock.Any()).Return(errors.New("error"))
				err := uc.CancelBooking(ctx, bookingId)
				So(err, ShouldNotBeNil)
			})
			Convey("Update Booking success", func() {
				mockBookingStorage.EXPECT().UpdateStatus(ctx, bookingId, gomock.Any()).Return(nil)
				err := uc.CancelBooking(ctx, bookingId)
				So(err, ShouldBeNil)
			})
		})
	})
}
