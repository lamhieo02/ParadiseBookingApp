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

func TestCreateCalendarGuider(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockCalendarGuiderSto := NewMockCalendarGuiderStorage(ctrl)
	// init usecase
	uc := NewCalendarGuiderUseCase(mockCalendarGuiderSto, nil)

	// data test
	dataTest := calendarguideriomodel.CreateCalendarGuiderReq{}
	if err := gofakeit.Struct(&dataTest); err != nil {
		t.Error("Error to create data test")
	}

	// 02-01-2006 15:04:05
	dataTest.DateFrom = "02-01-2024 15:04:05"
	dataTest.DateTo = "02-04-2024 15:04:05"

	Convey("Test Create Calendar Guider", t, func() {
		Convey("Create calendar guider fail", func() {
			mockCalendarGuiderSto.EXPECT().Create(ctx, gomock.Any()).Return(errors.New("error"))
			err := uc.CreateCalendarGuider(ctx, &dataTest)
			So(err, ShouldNotBeNil)
		})
		Convey("Create calendar guider success", func() {
			mockCalendarGuiderSto.EXPECT().Create(ctx, gomock.Any()).Return(nil)
			err := uc.CreateCalendarGuider(ctx, &dataTest)
			So(err, ShouldBeNil)
		})
	})

}
