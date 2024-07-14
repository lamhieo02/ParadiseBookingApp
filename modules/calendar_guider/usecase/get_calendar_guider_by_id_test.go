package calendarguiderusecase

import (
	"context"
	"errors"
	"paradise-booking/entities"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCalendarGuiderByID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockCalendarGuiderSto := NewMockCalendarGuiderStorage(ctrl)
	mockBookingGuiderSto := NewMockBookingGuiderSto(ctrl)
	// init usecase
	uc := NewCalendarGuiderUseCase(mockCalendarGuiderSto, mockBookingGuiderSto)

	id := gofakeit.Number(1, 100)

	Convey("Test Get Calendar Guider By ID", t, func() {
		Convey("Get calendar guider by ID fail", func() {
			mockCalendarGuiderSto.EXPECT().GetByID(ctx, id).Return(nil, errors.New("error"))
			_, err := uc.GetCalendarGuiderByID(ctx, id)
			So(err, ShouldNotBeNil)
		})
		Convey("Get calendar guider by ID success", func() {
			now := time.Now()
			mockCalendarGuiderSto.EXPECT().GetByID(ctx, id).Return(&entities.CalendarGuider{DateFrom: &now, DateTo: &now}, nil)
			Convey("List booking guider by condition", func() {
				mockBookingGuiderSto.EXPECT().ListByCondition(ctx, gomock.Any()).Return(nil, errors.New("error"))
				_, err := uc.GetCalendarGuiderByID(ctx, id)
				So(err, ShouldNotBeNil)
			})
			Convey("List booking guider by condition success", func() {
				resp := make([]*entities.BookingGuider, 3)
				for i := range resp {
					resp[i] = &entities.BookingGuider{}
				}
				mockBookingGuiderSto.EXPECT().ListByCondition(ctx, gomock.Any()).Return(resp, nil)
				_, err := uc.GetCalendarGuiderByID(ctx, id)
				So(err, ShouldBeNil)
			})
		})
	})

}
