package bookingusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	"paradise-booking/config"
	"paradise-booking/entities"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetBookingByPlaceID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockBookingStorage := NewMockBookingStorage(ctrl)
	mockBookingDetailStorage := NewMockBookingDetailStorage(ctrl)
	mockAccountSto := NewMockAccountSto(ctrl)
	mockPlaceSto := NewMockPlaceSto(ctrl)

	placeID := gofakeit.Number(1, 100)
	paging := common.Paging{}
	if err := gofakeit.Struct(&paging); err != nil {
		t.Error(err)
	}

	// init usecase
	cfg := config.Config{}
	uc := NewBookingUseCase(mockBookingStorage, mockBookingDetailStorage, &cfg, nil, mockAccountSto, mockPlaceSto, nil, nil)

	Convey("Test Get Booking By ID", t, func() {
		Convey("Get Booking By place ID fail", func() {
			mockBookingStorage.EXPECT().GetByPlaceID(ctx, placeID, &paging).Return(nil, errors.New("error"))
			_, err := uc.GetBookingByPlaceID(ctx, placeID, &paging)
			So(err, ShouldNotBeNil)
		})
		Convey("Get Booking By place ID success", func() {
			mockBookingStorage.EXPECT().GetByPlaceID(ctx, placeID, &paging).Return([]entities.Booking{}, nil)
			Convey("Get place by id fail", func() {
				mockPlaceSto.EXPECT().GetPlaceByID(ctx, placeID).Return(nil, errors.New("error"))
				_, err := uc.GetBookingByPlaceID(ctx, placeID, &paging)
				So(err, ShouldNotBeNil)
			})
			Convey("Get place by id success", func() {
				mockPlaceSto.EXPECT().GetPlaceByID(ctx, placeID).Return(&entities.Place{}, nil)
				_, err := uc.GetBookingByPlaceID(ctx, placeID, &paging)
				So(err, ShouldBeNil)
			})
		})
	})
}
