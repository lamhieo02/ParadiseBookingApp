package calendarguiderusecase

import (
	"context"
	"errors"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdateCalendarGuiderByID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockCalendarGuiderSto := NewMockCalendarGuiderStorage(ctrl)
	mockBookingGuiderSto := NewMockBookingGuiderSto(ctrl)

	// init usecase
	uc := NewCalendarGuiderUseCase(mockCalendarGuiderSto, mockBookingGuiderSto)

	id := gofakeit.Number(1, 100)
	dataTest := calendarguideriomodel.UpdateCalendarGuiderReq{}
	if err := gofakeit.Struct(&dataTest); err != nil {
		t.Error("Error to create data test")
	}

	dataTest.DateFrom = ""
	dataTest.DateTo = ""

	Convey("Test Update Calendar Guider By ID", t, func() {
		Convey("Update calendar guider fail", func() {
			mockCalendarGuiderSto.EXPECT().UpdateByID(ctx, id, gomock.Any()).Return(errors.New("error"))
			err := uc.UpdateCalendarGuiderByID(ctx, id, &dataTest)
			So(err, ShouldNotBeNil)
		})
		Convey("Update calendar guider success", func() {
			mockCalendarGuiderSto.EXPECT().UpdateByID(ctx, id, gomock.Any()).Return(nil)
			err := uc.UpdateCalendarGuiderByID(ctx, id, &dataTest)
			So(err, ShouldBeNil)
		})
	})
}
