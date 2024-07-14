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

func TestListPaymentByVendorID(t *testing.T) {
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

	Convey("Test List Payment By Vendor ID", t, func() {
		Convey("List payment by vendor ID fail", func() {
			mockPaymentSto.EXPECT().GetPaymentByVendor(ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			_, err := uc.ListPaymentByVendorID(ctx, &paging, guiderId, bookingId)
			So(err, ShouldNotBeNil)
		})
		Convey("List payment by guider ID success", func() {
			mockPaymentSto.EXPECT().GetPaymentByVendor(ctx, gomock.Any(), gomock.Any()).Return(nil, nil)
			_, err := uc.ListPaymentByVendorID(ctx, &paging, guiderId, bookingId)
			So(err, ShouldBeNil)
		})
	})
}
