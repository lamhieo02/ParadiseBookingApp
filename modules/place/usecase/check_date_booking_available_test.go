package placeusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"paradise-booking/entities"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckDateBookingAvailable(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockPlaceSto := NewMockPlaceStorage(ctrl)
	// mockPlaceCache := NewMockPlaceStoCache(ctrl)
	// mockAccountSto := NewMockAccountStorage(ctrl)
	// mockPlaceWishListSto := NewMockPlaceWishListSto(ctrl)
	mockBookingSto := NewMockBookingSto(ctrl)
	// mockPostGuideSto := NewMockPostGuideSto(ctrl)
	// mockPostGuideCache := NewMockPostGuideCache(ctrl)

	// init usecase
	cfg := config.Config{}
	uc := NewPlaceUseCase(&cfg, mockPlaceSto, nil, nil, nil, nil, mockBookingSto, nil, nil)
	// placeId int64, dateFrom string, dateTo string
	placeID := gofakeit.Number(1, 100)
	dateFrom := gofakeit.Date().Format("2006-01-02")
	dateTo := gofakeit.Date().Format("2006-01-02")

	Convey("Test Check Date Booking Available", t, func() {
		Convey("Get place by id fail", func() {
			mockPlaceSto.EXPECT().GetPlaceByID(ctx, placeID).Return(nil, errors.New("error"))
			_, err := uc.CheckDateBookingAvailable(ctx, int64(placeID), dateFrom, dateTo)
			So(err, ShouldNotBeNil)
		})
		Convey("Get place by id success", func() {
			mockPlaceSto.EXPECT().GetPlaceByID(ctx, placeID).Return(&entities.Place{}, nil)
			Convey("Get booking by place id fail", func() {
				mockBookingSto.EXPECT().GetBookingsWithinDateRange(ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
				_, err := uc.CheckDateBookingAvailable(ctx, int64(placeID), dateFrom, dateTo)
				So(err, ShouldNotBeNil)
			})
			Convey("Get booking by place id success", func() {
				res := []entities.Booking{}
				res = append(res, entities.Booking{})
				mockBookingSto.EXPECT().GetBookingsWithinDateRange(ctx, gomock.Any(), gomock.Any()).Return(res, nil)
				_, err := uc.CheckDateBookingAvailable(ctx, int64(placeID), dateFrom, dateTo)
				So(err, ShouldBeNil)
			})
		})
	})
}
