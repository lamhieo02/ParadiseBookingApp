package accountusecase

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

func TestUpdatePassword(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAccountStorage := NewMockAccountStorage(ctrl)

	// init usecase
	cfg := config.Config{}
	uc := NewUserUseCase(&cfg, mockAccountStorage, nil, nil, nil)

	// prepare data to test
	email := gofakeit.Email()
	newPassword := gofakeit.Password(true, true, true, true, false, 10)

	Convey("Test Update Password", t, func() {
		Convey("Get account by ID fail", func() {
			mockAccountStorage.EXPECT().GetAccountByEmail(ctx, email).Return(nil, errors.New("error"))
			err := uc.UpdatePassword(ctx, email, newPassword)
			So(err, ShouldNotBeNil)
		})
		Convey("Get account by ID success", func() {
			mockAccountStorage.EXPECT().GetAccountByEmail(ctx, email).Return(&entities.Account{}, nil)
			Convey("Update account fail", func() {
				mockAccountStorage.EXPECT().UpdateAccountById(ctx, gomock.Any(), gomock.Any()).Return(errors.New("error"))
				err := uc.UpdatePassword(ctx, email, newPassword)
				So(err, ShouldNotBeNil)
			})
			Convey("Update account success", func() {
				mockAccountStorage.EXPECT().UpdateAccountById(ctx, gomock.Any(), gomock.Any()).Return(nil)
				err := uc.UpdatePassword(ctx, email, newPassword)
				So(err, ShouldBeNil)
			})
		})
	})
}
