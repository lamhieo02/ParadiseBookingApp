package bookingratingusecase

import (
	"context"
	"errors"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCommentByBookingID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockBookingRatingSto := NewMockBookingRatingSto(ctrl)
	mockAccountSto := NewMockAccountSto(ctrl)
	mockPlaceSto := NewMockPlaceSto(ctrl)
	mockPostGuideSto := NewMockPostGuideSto(ctrl)
	// init usecase
	uc := Newbookingratingusecase(mockBookingRatingSto, mockAccountSto, mockPlaceSto, nil, mockPostGuideSto)

	bookingID := gofakeit.Number(1, 100)
	objectType := gofakeit.Number(1, 2)

	Convey("Test Get Comment By Booking ID", t, func() {
		Convey("Get Booking Rating by ID fail", func() {
			mockBookingRatingSto.EXPECT().GetByCondition(ctx, gomock.Any()).Return(nil, errors.New("error"))
			_, err := uc.GetCommentByBookingID(ctx, bookingID, objectType)
			So(err, ShouldNotBeNil)
		})
		Convey("Get Booking Rating by ID success", func() {
			bookingRatings := make([]entities.BookingRating, 4)
			for i := 0; i < 4; i++ {
				if err := gofakeit.Struct(&bookingRatings[i]); err != nil {
					t.Error(err)
				}
			}
			mockBookingRatingSto.EXPECT().GetByCondition(ctx, gomock.Any()).Return(bookingRatings, nil)
			Convey("Get Profile By ID success", func() {
				mockAccountSto.EXPECT().GetProfileByID(ctx, gomock.Any()).Return(&entities.Account{}, nil).AnyTimes()
				Convey("Object Type = Place", func() {
					objectType = constant.BookingRatingObjectTypePlace
					Convey("Get Place By ID success", func() {
						mockPlaceSto.EXPECT().GetPlaceByID(ctx, gomock.Any()).Return(&entities.Place{}, nil).AnyTimes()
						_, err := uc.GetCommentByBookingID(ctx, bookingID, objectType)
						So(err, ShouldBeNil)
					})
				})
				Convey("Object Type = Post Guide", func() {
					objectType = constant.BookingRatingObjectTypeGuide
					Convey("Get Place By ID success", func() {
						mockPostGuideSto.EXPECT().GetByID(ctx, gomock.Any()).Return(&entities.PostGuide{}, nil).AnyTimes()
						_, err := uc.GetCommentByBookingID(ctx, bookingID, objectType)
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})
}
