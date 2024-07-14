package paymentusecase

import (
	"context"
	"errors"
	"paradise-booking/entities"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdateStatusPaymentByID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockPaymentSto := NewMockPaymentSto(ctrl)

	// init usecase
	uc := NewPaymentUseCase(mockPaymentSto)

	id := gofakeit.Number(1, 100)
	status := gofakeit.Number(0, 1)

	Convey("Test Update Status Payment By ID", t, func() {
		Convey("Get payment by ID fail", func() {
			mockPaymentSto.EXPECT().GetByID(ctx, id).Return(nil, errors.New("error"))
			err := uc.UpdateStatusPaymentByID(ctx, id, status)
			So(err, ShouldNotBeNil)
		})
		Convey("Get payment by ID success", func() {
			mockPaymentSto.EXPECT().GetByID(ctx, id).Return(&entities.Payment{}, nil)
			Convey("Update status payment fail", func() {
				mockPaymentSto.EXPECT().UpdateByID(ctx, gomock.Any(), gomock.Any()).Return(errors.New("error"))
				err := uc.UpdateStatusPaymentByID(ctx, id, status)
				So(err, ShouldNotBeNil)
			})
			Convey("Update status payment success", func() {
				mockPaymentSto.EXPECT().UpdateByID(ctx, gomock.Any(), gomock.Any()).Return(nil)
				err := uc.UpdateStatusPaymentByID(ctx, id, status)
				So(err, ShouldBeNil)
			})
		})
	})
}
