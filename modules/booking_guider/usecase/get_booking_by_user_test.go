package bookingguiderusecase

import (
	"context"
	"errors"
	"paradise-booking/entities"
	"paradise-booking/worker"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetBookingByUserID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockBookingGuiderSto := NewMockbookingGuiderStorage(ctrl)
	mockPaymentSto := NewMockPaymentSto(ctrl)
	mockTaskDistributor := worker.NewMockTaskDistributor(ctrl)
	mockPostGuideSto := NewMockPostGuideUC(ctrl)
	mockCalendarSto := NewMockCalendarSto(ctrl)

	// init usecase
	uc := NewBookingGuiderUseCase(mockBookingGuiderSto, mockTaskDistributor, nil, mockPaymentSto, mockCalendarSto, mockPostGuideSto)

	userID := gofakeit.Number(1, 100)

	Convey("Test Get Booking By User ID", t, func() {
		Convey("List booking by user ID fail", func() {
			mockBookingGuiderSto.EXPECT().ListByCondition(ctx, gomock.Any()).Return(nil, errors.New("error"))
			_, err := uc.GetBookingByUserID(ctx, userID)
			So(err, ShouldNotBeNil)
		})
		Convey("List booking by user ID success", func() {
			bookingsGuider := make([]*entities.BookingGuider, 2)
			for i := 0; i < 2; i++ {
				bookingsGuider[i] = &entities.BookingGuider{}
				if err := gofakeit.Struct(bookingsGuider[i]); err != nil {
					t.Error(err)
				}
			}
			now := time.Now()
			mockBookingGuiderSto.EXPECT().ListByCondition(ctx, gomock.Any()).Return(bookingsGuider, nil)
			mockCalendarSto.EXPECT().GetByID(ctx, gomock.Any()).Return(&entities.CalendarGuider{DateFrom: &now, DateTo: &now}, nil).AnyTimes()
			_, err := uc.GetBookingByUserID(ctx, userID)
			So(err, ShouldBeNil)
		})
	})
}
