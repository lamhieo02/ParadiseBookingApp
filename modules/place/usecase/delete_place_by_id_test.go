package placeusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"paradise-booking/entities"
	googlemapprovider "paradise-booking/provider/googlemap"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDeletePlaceByID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockPlaceSto := NewMockPlaceStorage(ctrl)
	// mockPlaceCache := NewMockPlaceStoCache(ctrl)
	mockAccountSto := NewMockAccountStorage(ctrl)
	// mockPlaceWishListSto := NewMockPlaceWishListSto(ctrl)
	// mockBookingSto := NewMockBookingSto(ctrl)
	// mockPostGuideSto := NewMockPostGuideSto(ctrl)
	// mockPostGuideCache := NewMockPostGuideCache(ctrl)
	cfg := config.Config{}
	mockGoogleMap := googlemapprovider.NewGoogleMap(&cfg)
	uc := NewPlaceUseCase(&cfg, mockPlaceSto, mockAccountSto, mockGoogleMap, nil, nil, nil, nil, nil)

	placeID := gofakeit.Number(1, 100)
	email := gofakeit.Email()
	Convey("Test Delete Place By ID", t, func() {
		Convey("Get Place By ID", func() {
			mockPlaceSto.EXPECT().GetPlaceByID(ctx, gomock.Any()).Return(nil, errors.New("error"))
			err := uc.DeletePlaceByID(ctx, placeID, email)
			So(err, ShouldNotBeNil)
		})
		Convey("Get Place By ID Success", func() {
			mockPlaceSto.EXPECT().GetPlaceByID(ctx, gomock.Any()).Return(&entities.Place{}, nil)
			Convey("Get Account By Email", func() {
				mockAccountSto.EXPECT().GetAccountByEmail(ctx, gomock.Any()).Return(nil, errors.New("error"))
				err := uc.DeletePlaceByID(ctx, placeID, email)
				So(err, ShouldNotBeNil)
			})
			Convey("Get Account By Email Success", func() {
				mockAccountSto.EXPECT().GetAccountByEmail(ctx, gomock.Any()).Return(&entities.Account{}, nil)
				Convey("Delete Place Fail", func() {
					mockPlaceSto.EXPECT().DeleteByID(ctx, gomock.Any()).Return(errors.New("error"))
					err := uc.DeletePlaceByID(ctx, placeID, email)
					So(err, ShouldNotBeNil)
				})
				Convey("Delete Place Success", func() {
					mockPlaceSto.EXPECT().DeleteByID(ctx, gomock.Any()).Return(nil)
					err := uc.DeletePlaceByID(ctx, placeID, email)
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
