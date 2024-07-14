package bookingusecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListBookingByCondition(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockBookingStorage := NewMockBookingStorage(ctrl)

	// init usecase
	uc := NewBookingUseCase(mockBookingStorage, nil, nil, nil, nil, nil, nil, nil)

	Convey("Test List Booking By Condition", t, func() {
		Convey("List Booking By Condition fail", func() {
			mockBookingStorage.EXPECT().ListAllBookingWithCondition(ctx, gomock.Any()).Return(nil, errors.New("error"))
			_, err := uc.ListBookingByCondition(ctx)
			So(err, ShouldNotBeNil)
		})
		Convey("List Booking By Condition success", func() {
			mockBookingStorage.EXPECT().ListAllBookingWithCondition(ctx, gomock.Any()).Return(nil, nil)
			_, err := uc.ListBookingByCondition(ctx)
			So(err, ShouldBeNil)
		})
	})
}
