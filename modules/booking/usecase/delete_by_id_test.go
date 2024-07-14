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

func TestDeleteBookingByID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockBookingStorage := NewMockBookingStorage(ctrl)

	// init usecase
	cfg := config.Config{}
	uc := NewBookingUseCase(mockBookingStorage, nil, &cfg, nil, nil, nil, nil, nil)

	bookingID := gofakeit.Number(1, 100)

	Convey("Test Delete Booking By ID", t, func() {
		Convey("Get Booking by id fail", func() {
			mockBookingStorage.EXPECT().GetByID(ctx, bookingID).Return(nil, errors.New("error"))
			err := uc.DeleteBookingByID(ctx, bookingID)
			So(err, ShouldNotBeNil)
		})
		Convey("Get Booking by id success", func() {
			mockBookingStorage.EXPECT().GetByID(ctx, bookingID).Return(&entities.Booking{StatusId: constant.BookingStatusCompleted}, nil)
			Convey("Delete Booking fail", func() {
				mockBookingStorage.EXPECT().DeleteByID(ctx, bookingID).Return(errors.New("error"))
				err := uc.DeleteBookingByID(ctx, bookingID)
				So(err, ShouldNotBeNil)
			})
			Convey("Delete Booking success", func() {
				mockBookingStorage.EXPECT().DeleteByID(ctx, bookingID).Return(nil)
				err := uc.DeleteBookingByID(ctx, bookingID)
				So(err, ShouldBeNil)
			})
		})
	})
}
