package bookingratingusecase

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetStatisticByObjectID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockBookingRatingSto := NewMockBookingRatingSto(ctrl)
	mockAccountSto := NewMockAccountSto(ctrl)
	mockPlaceSto := NewMockPlaceSto(ctrl)
	mockPostGuideSto := NewMockPostGuideSto(ctrl)
	// init usecase
	uc := Newbookingratingusecase(mockBookingRatingSto, mockAccountSto, mockPlaceSto, nil, mockPostGuideSto)

	objectID := gofakeit.Number(1, 100)
	objectType := gofakeit.Number(1, 2)

	Convey("Test Get Statistic By Object ID", t, func() {
		Convey("Get Statistic By Object ID fail", func() {
			mockBookingRatingSto.EXPECT().GetStatisticByObjectID(ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			_, err := uc.GetStatisticByObjectID(ctx, objectID, objectType)
			So(err, ShouldNotBeNil)
		})
		Convey("Get Statistic By Object ID success", func() {
			mockBookingRatingSto.EXPECT().GetStatisticByObjectID(ctx, gomock.Any(), gomock.Any()).Return(nil, nil)
			_, err := uc.GetStatisticByObjectID(ctx, objectID, objectType)
			So(err, ShouldBeNil)

		})
	})
}
