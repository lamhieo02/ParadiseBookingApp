package bookingguiderusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	"paradise-booking/entities"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
	"paradise-booking/worker"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListBooking(t *testing.T) {
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

	paging := common.Paging{}
	if err := gofakeit.Struct(&paging); err != nil {
		t.Error(err)

	}
	filter := bookingguideriomodel.Filter{}
	if err := gofakeit.Struct(&filter); err != nil {
		t.Error(err)
	}

	Convey("Test List Booking", t, func() {
		Convey("List by filter fail", func() {
			mockBookingGuiderSto.EXPECT().ListByFilter(ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			_, err := uc.ListBooking(ctx, &paging, &filter)
			So(err, ShouldNotBeNil)
		})
		Convey("List by filter success", func() {
			data := make([]entities.BookingGuider, 2)
			now := time.Now()
			for i := 0; i < 2; i++ {
				data[i] = entities.BookingGuider{}
				if err := gofakeit.Struct(&data[i]); err != nil {
					t.Error(err)
				}
			}
			mockBookingGuiderSto.EXPECT().ListByFilter(ctx, gomock.Any(), gomock.Any()).Return(data, nil)
			mockPostGuideSto.EXPECT().GetPostGuideByID(ctx, gomock.Any()).Return(&postguideiomodel.GetPostGuideResp{}, nil).AnyTimes()
			mockCalendarSto.EXPECT().GetByID(ctx, gomock.Any()).Return(&entities.CalendarGuider{DateFrom: &now, DateTo: &now}, nil).AnyTimes()
			_, err := uc.ListBooking(ctx, &paging, &filter)
			So(err, ShouldBeNil)
		})
	})
}
