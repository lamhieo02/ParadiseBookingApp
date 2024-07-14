package bookingusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"paradise-booking/entities"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetBookingByID(t *testing.T) {
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

	bookingID := gofakeit.Number(1, 100)

	Convey("Test Get Booking By ID", t, func() {
		Convey("Get Booking By ID fail", func() {
			mockBookingStorage.EXPECT().GetByID(ctx, bookingID).Return(nil, nil)
			_, err := uc.GetBookingByID(ctx, bookingID)
			So(err, ShouldNotBeNil)
		})
		Convey("Get Booking By ID success", func() {
			booking := &entities.Booking{}
			booking.Id = bookingID
			now := time.Now()
			booking.CheckInDate = &now
			booking.ChekoutDate = &now
			mockBookingStorage.EXPECT().GetByID(ctx, bookingID).Return(booking, nil)
			Convey("Get Booking Detail fail", func() {
				mockBookingDetailStorage.EXPECT().GetByBookingID(ctx, bookingID).Return(nil, errors.New("error"))
				_, err := uc.GetBookingByID(ctx, bookingID)
				So(err, ShouldNotBeNil)
			})
			Convey("Get Booking Detail success", func() {
				mockBookingDetailStorage.EXPECT().GetByBookingID(ctx, bookingID).Return(&entities.BookingDetail{}, nil)
				Convey("Get account by id fail", func() {
					mockAccountSto.EXPECT().GetProfileByID(ctx, gomock.Any()).Return(nil, errors.New("error"))
					_, err := uc.GetBookingByID(ctx, bookingID)
					So(err, ShouldNotBeNil)
				})
				Convey("Get account by id success", func() {
					mockAccountSto.EXPECT().GetProfileByID(ctx, gomock.Any()).Return(&entities.Account{}, nil)
					Convey("Get place by id fail", func() {
						mockPlaceSto.EXPECT().GetPlaceByID(ctx, gomock.Any()).Return(nil, errors.New("error"))
						_, err := uc.GetBookingByID(ctx, bookingID)
						So(err, ShouldNotBeNil)
					})
					Convey("Get place by id success", func() {
						mockPlaceSto.EXPECT().GetPlaceByID(ctx, gomock.Any()).Return(&entities.Place{}, nil)
						_, err := uc.GetBookingByID(ctx, bookingID)
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})
}
