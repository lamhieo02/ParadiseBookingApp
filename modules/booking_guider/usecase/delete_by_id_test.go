package bookingguiderusecase

import (
	"context"
	"errors"
	"paradise-booking/worker"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDeleteBookingByID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockBookingGuiderSto := NewMockbookingGuiderStorage(ctrl)
	mockPaymentSto := NewMockPaymentSto(ctrl)
	mockTaskDistributor := worker.NewMockTaskDistributor(ctrl)

	// init usecase
	uc := NewBookingGuiderUseCase(mockBookingGuiderSto, mockTaskDistributor, nil, mockPaymentSto, nil, nil)

	id := gofakeit.Number(1, 100)

	Convey("Test Delete Booking By ID", t, func() {
		Convey("Delete booking fail", func() {
			mockBookingGuiderSto.EXPECT().DeleteByID(ctx, id).Return(errors.New("error"))
			err := uc.DeleteBookingByID(ctx, id)
			So(err, ShouldNotBeNil)
		})
		Convey("Delete booking success", func() {
			mockBookingGuiderSto.EXPECT().DeleteByID(ctx, id).Return(nil)
			err := uc.DeleteBookingByID(ctx, id)
			So(err, ShouldBeNil)
		})
	})
}
