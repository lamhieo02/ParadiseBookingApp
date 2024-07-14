package bookingusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	"paradise-booking/config"
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListBooking(t *testing.T) {
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

	// data test
	userID := gofakeit.Number(1, 100)
	paging := common.Paging{}
	if err := gofakeit.Struct(&paging); err != nil {
		t.Error(err)
	}
	filter := iomodel.FilterListBooking{}
	if err := gofakeit.Struct(&filter); err != nil {
		t.Error(err)
	}

	Convey("Test ListBooking", t, func() {
		Convey("List booking by filter fail", func() {
			mockBookingStorage.EXPECT().ListByFilter(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			_, err := uc.ListBooking(ctx, &paging, &filter, userID)
			So(err, ShouldNotBeNil)
		})
		Convey("List booking by filter success", func() {
			mockBookingStorage.EXPECT().ListByFilter(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]entities.Booking{}, nil)
			Convey("Get account by id fail", func() {
				mockAccountSto.EXPECT().GetProfileByID(ctx, userID).Return(nil, errors.New("error"))
				_, err := uc.ListBooking(ctx, &paging, &filter, userID)
				So(err, ShouldNotBeNil)
			})
			Convey("Get account by id success", func() {
				mockAccountSto.EXPECT().GetProfileByID(ctx, userID).Return(&entities.Account{}, nil)
				_, err := uc.ListBooking(ctx, &paging, &filter, userID)
				So(err, ShouldBeNil)
			})
		})
	})
}
