package paymentusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListPaymentByGuiderID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockPaymentSto := NewMockPaymentSto(ctrl)

	// init usecase
	uc := NewPaymentUseCase(mockPaymentSto)

	paging := common.Paging{}
	if err := gofakeit.Struct(&paging); err != nil {
		t.Error("Error to create data test")
	}
	paging.Limit = 0
	guiderId := gofakeit.Number(1, 100)
	bookingId := gofakeit.Number(1, 100)

	Convey("Test List Payment By Guider ID", t, func() {
		Convey("List payment by guider ID fail", func() {
			mockPaymentSto.EXPECT().GetPaymentByGuider(ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			_, err := uc.ListPaymentByGuiderID(ctx, &paging, guiderId, bookingId)
			So(err, ShouldNotBeNil)
		})
		Convey("List payment by guider ID success", func() {
			mockPaymentSto.EXPECT().GetPaymentByGuider(ctx, gomock.Any(), gomock.Any()).Return(nil, nil)
			_, err := uc.ListPaymentByGuiderID(ctx, &paging, guiderId, bookingId)
			So(err, ShouldBeNil)
		})
	})
}
