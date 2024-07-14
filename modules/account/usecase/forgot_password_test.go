package accountusecase

import (
	"context"
	"errors"
	"paradise-booking/entities"
	"paradise-booking/worker"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

func TestForgotPassword(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAccountStorage := NewMockAccountStorage(ctrl)
	mockTaskDistributor := worker.NewMockTaskDistributor(ctrl)

	// init usecase
	uc := NewUserUseCase(nil, mockAccountStorage, nil, mockTaskDistributor, nil)

	// prepare data to test
	email := gofakeit.Email()
	Convey("Test Forgot Password", t, func() {
		Convey("Check email is not exist", func() {
			mockAccountStorage.EXPECT().GetAccountByEmail(ctx, email).Return(nil, gorm.ErrRecordNotFound)
			err := uc.ForgotPassword(ctx, email)
			So(err, ShouldNotBeNil)
		})
		Convey("Check email is exist", func() {
			mockAccountStorage.EXPECT().GetAccountByEmail(ctx, email).Return(&entities.Account{}, nil)
			Convey("Send email fail", func() {
				mockTaskDistributor.EXPECT().DistributeTaskSendVerifyResetCodePassword(ctx, gomock.Any(), gomock.Any()).Return(errors.New("error"))
				err := uc.ForgotPassword(ctx, email)
				So(err, ShouldBeNil)
			})
			Convey("Send email success", func() {
				mockTaskDistributor.EXPECT().DistributeTaskSendVerifyResetCodePassword(ctx, gomock.Any(), gomock.Any()).Return(nil)
				err := uc.ForgotPassword(ctx, email)
				So(err, ShouldBeNil)
			})
		})
	})
}
