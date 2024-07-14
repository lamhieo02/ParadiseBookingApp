package bookingguiderusecase

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdateStatusBooking(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockBookingGuiderSto := NewMockbookingGuiderStorage(ctrl)

	// init usecase
	uc := NewBookingGuiderUseCase(mockBookingGuiderSto, nil, nil, nil, nil, nil)

	bookingGuiderID := gofakeit.Number(1, 100)
	status := gofakeit.Number(0, 4)

	Convey("Test Update Status Booking", t, func() {
		Convey("Update status booking fail", func() {
			mockBookingGuiderSto.EXPECT().UpdateStatus(ctx, bookingGuiderID, status).Return(errors.New("error"))
			err := uc.UpdateStatusBooking(ctx, bookingGuiderID, status)
			So(err, ShouldNotBeNil)
		})
		Convey("Update status booking success", func() {
			mockBookingGuiderSto.EXPECT().UpdateStatus(ctx, bookingGuiderID, status).Return(nil)
			err := uc.UpdateStatusBooking(ctx, bookingGuiderID, status)
			So(err, ShouldBeNil)
		})
	})
}
