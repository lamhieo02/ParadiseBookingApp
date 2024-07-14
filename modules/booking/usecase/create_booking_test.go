package bookingusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"paradise-booking/modules/booking/iomodel"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateBooking(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockBookingStorage := NewMockBookingStorage(ctrl)
	mockBookingDetailStorage := NewMockBookingDetailStorage(ctrl)
	paymentSto := NewMockPaymentSto(ctrl)

	// init usecase
	cfg := config.Config{}
	uc := NewBookingUseCase(mockBookingStorage, mockBookingDetailStorage, &cfg, nil, nil, nil, nil, paymentSto)

	// data test
	bookingData := iomodel.CreateBookingReq{}
	err := gofakeit.Struct(&bookingData)
	if err != nil {
		t.Error(err)
	}

	Convey("Test Create Booking", t, func() {
		Convey("Create Booking Fail", func() {
			mockBookingStorage.EXPECT().Create(ctx, gomock.Any()).Return(errors.New("error"))
			_, err := uc.CreateBooking(ctx, &bookingData)
			So(err, ShouldNotBeNil)
		})
		Convey("Create Booking Success", func() {
			Convey("Create booking detail fail", func() {
				mockBookingStorage.EXPECT().Create(ctx, gomock.Any()).Return(nil)
				mockBookingDetailStorage.EXPECT().CreateTx(ctx, gomock.Any()).Return(errors.New("error"))
				_, err := uc.CreateBooking(ctx, &bookingData)
				So(err, ShouldNotBeNil)
				Convey("Create Payment Fail", func() {
					mockBookingStorage.EXPECT().Create(ctx, gomock.Any()).Return(nil)
					mockBookingDetailStorage.EXPECT().CreateTx(ctx, gomock.Any()).Return(nil)
					paymentSto.EXPECT().CreatePayment(ctx, gomock.Any()).Return(errors.New("error"))
					_, err := uc.CreateBooking(ctx, &bookingData)
					So(err, ShouldNotBeNil)
				})
				Convey("Create Payment Success", func() {
					mockBookingStorage.EXPECT().Create(ctx, gomock.Any()).Return(nil)
					mockBookingDetailStorage.EXPECT().CreateTx(ctx, gomock.Any()).Return(nil)
					paymentSto.EXPECT().CreatePayment(ctx, gomock.Any()).Return(nil)
					_, err := uc.CreateBooking(ctx, &bookingData)
					So(err, ShouldBeNil)
				})
			})
		})
	})

}
