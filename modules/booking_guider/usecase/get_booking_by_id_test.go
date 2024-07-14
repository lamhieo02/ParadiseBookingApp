package bookingguiderusecase

import (
	"context"
	"errors"
	"paradise-booking/entities"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
	"paradise-booking/worker"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetBookingByID(t *testing.T) {
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

	bookingGuiderID := gofakeit.Number(1, 100)

	Convey("Test Get Booking By ID", t, func() {
		Convey("Get booking by ID fail", func() {
			mockBookingGuiderSto.EXPECT().GetByID(ctx, bookingGuiderID).Return(nil, errors.New("error"))
			_, err := uc.GetBookingByID(ctx, bookingGuiderID)
			So(err, ShouldNotBeNil)
		})
		Convey("Get booking by ID success", func() {
			mockBookingGuiderSto.EXPECT().GetByID(ctx, bookingGuiderID).Return(&entities.BookingGuider{}, nil)
			Convey("Get Post Guider and calendar By ID fail", func() {
				mockPostGuideSto.EXPECT().GetPostGuideByID(ctx, gomock.Any()).Return(nil, errors.New("error"))
				mockCalendarSto.EXPECT().GetByID(ctx, gomock.Any()).Return(nil, errors.New("error"))
				_, err := uc.GetBookingByID(ctx, bookingGuiderID)
				So(err, ShouldNotBeNil)
			})
			Convey("Get Post Guider and calendar By ID success", func() {
				now := time.Now()
				mockPostGuideSto.EXPECT().GetPostGuideByID(ctx, gomock.Any()).Return(&postguideiomodel.GetPostGuideResp{}, nil)
				mockCalendarSto.EXPECT().GetByID(ctx, gomock.Any()).Return(&entities.CalendarGuider{DateFrom: &now, DateTo: &now}, nil)
				_, err := uc.GetBookingByID(ctx, bookingGuiderID)
				So(err, ShouldBeNil)
			})
		})
	})
}
