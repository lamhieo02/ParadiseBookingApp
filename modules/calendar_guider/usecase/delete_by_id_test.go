package calendarguiderusecase

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDeleteCalendarGuiderByID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockCalendarGuiderSto := NewMockCalendarGuiderStorage(ctrl)
	// init usecase
	uc := NewCalendarGuiderUseCase(mockCalendarGuiderSto, nil)

	id := gofakeit.Number(1, 100)

	Convey("Test Delete Calendar Guider By ID", t, func() {
		Convey("Delete calendar guider fail", func() {
			mockCalendarGuiderSto.EXPECT().DeleteByID(ctx, id).Return(errors.New("error"))
			err := uc.DeleteCalendarGuiderByID(ctx, id)
			So(err, ShouldNotBeNil)
		})
		Convey("Delete calendar guider success", func() {
			mockCalendarGuiderSto.EXPECT().DeleteByID(ctx, id).Return(nil)
			err := uc.DeleteCalendarGuiderByID(ctx, id)
			So(err, ShouldBeNil)
		})
	})
}
