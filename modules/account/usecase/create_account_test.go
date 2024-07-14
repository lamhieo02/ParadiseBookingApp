package accountusecase

import (
	"context"
	"errors"
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateAccount(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAccountStorage := NewMockAccountStorage(ctrl)
	mockTaskDistributor := NewMockVerifyEmailsUseCase(ctrl)

	// init usecase
	uc := NewUserUseCase(nil, mockAccountStorage, mockTaskDistributor, nil, nil)

	// prepare data to test
	dataTests := make([]iomodel.AccountRegister, 10)
	for i := 0; i < 10; i++ {
		dataTests[i] = iomodel.AccountRegister{
			Email:    gofakeit.Email(),
			Password: gofakeit.Password(true, true, true, true, false, 10),
		}
	}

	for _, tc := range dataTests {
		Convey("Test Create Account", t, func() {
			Convey("Check email is existed", func() {
				mockAccountStorage.EXPECT().GetAccountByEmail(ctx, tc.Email).Return(&entities.Account{}, nil)

				result, err := uc.CreateAccount(ctx, &tc)
				convey.So(result, convey.ShouldBeNil)
				convey.So(err, convey.ShouldNotBeNil)
				convey.So(err.Error(), convey.ShouldEqual, "email is existed")
			})
			Convey("Check email is not exist", func() {
				mockAccountStorage.EXPECT().GetAccountByEmail(ctx, tc.Email).Return(nil, nil)
				Convey("Create account fail", func() {
					mockAccountStorage.EXPECT().CreateTx(ctx, gomock.Any()).Return(errors.New("error"))
					result, err := uc.CreateAccount(ctx, &tc)
					convey.So(result, convey.ShouldBeNil)
					convey.So(err, convey.ShouldNotBeNil)
				})
				Convey("Create account success", func() {
					mockAccountStorage.EXPECT().CreateTx(ctx, gomock.Any()).Return(nil)
					result, err := uc.CreateAccount(ctx, &tc)
					convey.So(result, convey.ShouldNotBeNil)
					convey.So(err, convey.ShouldBeNil)
				})
			})
		})
	}
}
