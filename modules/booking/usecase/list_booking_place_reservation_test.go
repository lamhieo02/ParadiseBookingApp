package bookingusecase

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

func TestListPlaceReservationByVendor(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockBookingStorage := NewMockBookingStorage(ctrl)
	mockBookingDetailStorage := NewMockBookingDetailStorage(ctrl)
	mockAccountSto := NewMockAccountSto(ctrl)
	mockPlaceSto := NewMockPlaceSto(ctrl)

	// init usecase
	cfg := config.Config{}
	uc := NewBookingUseCase(mockBookingStorage, mockBookingDetailStorage, &cfg, nil, mockAccountSto, mockPlaceSto, nil, nil)

	vendorID := gofakeit.Number(1, 100)
	placeID := gofakeit.Number(1, 100)

	Convey("Test List Place Reservation By Vendor", t, func() {
		Convey("Place id != 0", func() {
			Convey("Get Place By ID fail", func() {
				mockPlaceSto.EXPECT().GetPlaceByID(ctx, placeID).Return(nil, errors.New("error"))
				_, err := uc.ListPlaceReservationByVendor(ctx, vendorID, placeID)
				So(err, ShouldNotBeNil)
			})
			Convey("Get Place By ID success", func() {
				mockPlaceSto.EXPECT().GetPlaceByID(ctx, placeID).Return(&entities.Place{}, nil)
				Convey("List all booking with condition fail", func() {
					mockBookingStorage.EXPECT().ListAllBookingWithCondition(ctx, gomock.Any()).Return(nil, errors.New("error"))
					_, err := uc.ListPlaceReservationByVendor(ctx, vendorID, placeID)
					So(err, ShouldNotBeNil)
				})
				Convey("List all booking with condition success", func() {
					mockBookingStorage.EXPECT().ListAllBookingWithCondition(ctx, gomock.Any()).Return([]entities.Booking{}, nil)
					_, err := uc.ListPlaceReservationByVendor(ctx, vendorID, placeID)
					So(err, ShouldBeNil)
				})
			})
		})
		Convey("Place id == 0", func() {
			placeID = 0
			Convey("List place with condition fail", func() {
				mockPlaceSto.EXPECT().ListPlaceByCondition(ctx, gomock.Any()).Return(nil, errors.New("error"))
				_, err := uc.ListPlaceReservationByVendor(ctx, vendorID, placeID)
				So(err, ShouldNotBeNil)
			})
			Convey("List place with condition success", func() {
				places := make([]entities.Place, 2)
				for i := 0; i < 2; i++ {
					if err := gofakeit.Struct(&places[i]); err != nil {
						t.Error(err)
					}
				}
				mockPlaceSto.EXPECT().ListPlaceByCondition(ctx, gomock.Any()).Return(places, nil)
				Convey("List all booking with condition fail", func() {
					mockBookingStorage.EXPECT().ListAllBookingWithCondition(ctx, gomock.Any()).Return(nil, errors.New("error"))
					_, err := uc.ListPlaceReservationByVendor(ctx, vendorID, placeID)
					So(err, ShouldNotBeNil)
				})
				Convey("List all booking with condition success", func() {
					mockBookingStorage.EXPECT().ListAllBookingWithCondition(ctx, gomock.Any()).Return([]entities.Booking{}, nil)
					_, err := uc.ListPlaceReservationByVendor(ctx, vendorID, placeID)
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
