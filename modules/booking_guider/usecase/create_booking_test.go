package bookingguiderusecase

import (
	"context"
	"errors"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
	"paradise-booking/worker"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateBookingGuider(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockBookingGuiderSto := NewMockbookingGuiderStorage(ctrl)
	mockPaymentSto := NewMockPaymentSto(ctrl)
	mockTaskDistributor := worker.NewMockTaskDistributor(ctrl)

	// init usecase
	uc := NewBookingGuiderUseCase(mockBookingGuiderSto, mockTaskDistributor, nil, mockPaymentSto, nil, nil)

	// prepare data to test
	dataTest := bookingguideriomodel.CreateBookingReq{}
	if err := gofakeit.Struct(&dataTest); err != nil {
		t.Error(err)
	}

	dataTest.PaymentMethod = 1

	Convey("Test Create Booking Guider", t, func() {
		Convey("Create booking guider fail", func() {
			mockBookingGuiderSto.EXPECT().Create(ctx, gomock.Any()).Return(errors.New("error"))
			if _, err := uc.CreateBookingGuider(ctx, &dataTest); err != nil {
				So(err, ShouldNotBeNil)
			}
		})
		Convey("Create booking guider success", func() {
			mockBookingGuiderSto.EXPECT().Create(ctx, gomock.Any()).Return(nil)
			mockTaskDistributor.EXPECT().DistributeTaskSendConfirmBookingGuider(ctx, gomock.Any(), gomock.Any()).Return(nil)
			Convey("Create payment fail", func() {
				mockPaymentSto.EXPECT().CreatePayment(ctx, gomock.Any()).Return(errors.New("error"))
				if _, err := uc.CreateBookingGuider(ctx, &dataTest); err != nil {
					So(err, ShouldNotBeNil)
				}
			})
			Convey("Create payment success", func() {
				mockPaymentSto.EXPECT().CreatePayment(ctx, gomock.Any()).Return(nil)
				_, err := uc.CreateBookingGuider(ctx, &dataTest)
				So(err, ShouldBeNil)

			})
		})
	})
}
