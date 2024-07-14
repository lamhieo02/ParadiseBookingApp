package bookingusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdateStatusBooking(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockBookingStorage := NewMockBookingStorage(ctrl)

	bookingID := gofakeit.Number(1, 100)
	status := gofakeit.Number(1, 4)
	// init usecase
	cfg := config.Config{}
	uc := NewBookingUseCase(mockBookingStorage, nil, &cfg, nil, nil, nil, nil, nil)

	Convey("Test Update Status Booking", t, func() {
		Convey("Update Status Booking fail", func() {
			mockBookingStorage.EXPECT().UpdateStatus(ctx, bookingID, status).Return(errors.New("error"))
			err := uc.UpdateStatusBooking(ctx, bookingID, status)
			So(err, ShouldNotBeNil)
		})
		Convey("Update Status Booking success", func() {
			mockBookingStorage.EXPECT().UpdateStatus(ctx, bookingID, status).Return(nil)
			err := uc.UpdateStatusBooking(ctx, bookingID, status)
			So(err, ShouldBeNil)
		})
	})

}
