package calendarguiderusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListCalendarGuiderByFilter(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockCalendarGuiderSto := NewMockCalendarGuiderStorage(ctrl)
	mockBookingGuiderSto := NewMockBookingGuiderSto(ctrl)

	// init usecase
	uc := NewCalendarGuiderUseCase(mockCalendarGuiderSto, mockBookingGuiderSto)

	// data test
	paging := common.Paging{}
	filter := calendarguideriomodel.Filter{}

	Convey("Test List Calendar Guider By Filter", t, func() {
		Convey("List calendar guider and booking guider by filter fail", func() {
			mockCalendarGuiderSto.EXPECT().ListByFilter(ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			mockBookingGuiderSto.EXPECT().ListByCondition(ctx, gomock.Any()).Return(nil, errors.New("error"))
			_, err := uc.ListCalendarGuiderByFilter(ctx, &paging, &filter)
			So(err, ShouldNotBeNil)
		})
		Convey("List calendar guider and booking guider by filter success", func() {
			mockCalendarGuiderSto.EXPECT().ListByFilter(ctx, gomock.Any(), gomock.Any()).Return(nil, nil)
			mockBookingGuiderSto.EXPECT().ListByCondition(ctx, gomock.Any()).Return(nil, nil)
			_, err := uc.ListCalendarGuiderByFilter(ctx, &paging, &filter)
			So(err, ShouldBeNil)
		})
	})
}
