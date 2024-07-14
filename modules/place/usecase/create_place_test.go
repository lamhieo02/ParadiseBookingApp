package placeusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"paradise-booking/entities"
	"paradise-booking/modules/place/iomodel"
	googlemapprovider "paradise-booking/provider/googlemap"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreatePlace(t *testing.T) {
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

	// data test
	data := iomodel.CreatePlaceReq{}
	email := gofakeit.Email()
	Convey("Test Create Place", t, func() {
		Convey("GetAccountByEmail fail", func() {
			mockAccountSto.EXPECT().GetAccountByEmail(ctx, gomock.Any()).Return(nil, errors.New("error"))
			err := uc.CreatePlace(ctx, &data, email)
			So(err, ShouldNotBeNil)
		})
		Convey("GetAccountByEmail success", func() {
			mockAccountSto.EXPECT().GetAccountByEmail(ctx, gomock.Any()).Return(&entities.Account{}, nil)
			Convey("CreatePlace fail", func() {
				mockPlaceSto.EXPECT().Create(ctx, gomock.Any()).Return(errors.New("error"))
				err := uc.CreatePlace(ctx, &data, email)
				So(err, ShouldNotBeNil)
			})
			Convey("CreatePlace success", func() {
				mockPlaceSto.EXPECT().Create(ctx, gomock.Any()).Return(nil)
				err := uc.CreatePlace(ctx, &data, email)
				So(err, ShouldBeNil)
			})
		})
	})
}
