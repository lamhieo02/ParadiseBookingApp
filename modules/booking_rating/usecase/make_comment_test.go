package bookingratingusecase

import (
	"context"
	"errors"
	"paradise-booking/constant"
	bookingratingiomodel "paradise-booking/modules/booking_rating/iomodel"
	"paradise-booking/provider/cache"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeComment(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockBookingRatingSto := NewMockBookingRatingSto(ctrl)
	mockAccountSto := NewMockAccountSto(ctrl)
	mockPlaceSto := NewMockPlaceSto(ctrl)
	mockPostGuideSto := NewMockPostGuideSto(ctrl)
	mockCache := cache.NewMockCache(ctrl)
	// init usecase
	uc := Newbookingratingusecase(mockBookingRatingSto, mockAccountSto, mockPlaceSto, mockCache, mockPostGuideSto)

	userID := gofakeit.Number(1, 100)
	testCase := make([]bookingratingiomodel.CreateBookingRatingReq, 2)
	for i := 0; i < 2; i++ {
		dataTest := bookingratingiomodel.CreateBookingRatingReq{}
		if err := gofakeit.Struct(&dataTest); err != nil {
			t.Error(err)
		}

		testCase[i] = dataTest
	}

	testCase[0].ObjectType = constant.BookingRatingObjectTypePlace
	testCase[1].ObjectType = constant.BookingRatingObjectTypeGuide

	for _, dataTest := range testCase {
		Convey("Test Make Comment", t, func() {
			Convey("Make Comment fail", func() {
				mockBookingRatingSto.EXPECT().Create(ctx, gomock.Any()).Return(nil, errors.New("error"))
				_, err := uc.MakeComment(ctx, userID, &dataTest)
				So(err, ShouldNotBeNil)
			})
			Convey("Make Comment success", func() {
				mockCache.EXPECT().Delete(gomock.Any(), gomock.Any())
				mockBookingRatingSto.EXPECT().Create(ctx, gomock.Any()).Return(nil, nil)
				_, err := uc.MakeComment(ctx, userID, &dataTest)
				So(err, ShouldBeNil)
			})
		})
	}
}
