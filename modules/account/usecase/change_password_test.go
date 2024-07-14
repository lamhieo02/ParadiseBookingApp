package accountusecase

import (
	"context"
	"errors"
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

func TestChangePassword(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAccountStorage := NewMockAccountStorage(ctrl)

	// init usecase
	uc := NewUserUseCase(nil, mockAccountStorage, nil, nil, nil)

	// prepare data to test
	email := gofakeit.Email()
	changePassModel := iomodel.ChangePassword{}
	changePassModel.OldPassword = gofakeit.Password(true, true, true, true, false, 10)
	changePassModel.NewPassword = gofakeit.Password(true, true, true, true, false, 10)

	Convey("Test Change Password", t, func() {
		Convey("Check email is not exist", func() {
			mockAccountStorage.EXPECT().GetAccountByEmail(ctx, email).Return(nil, gorm.ErrRecordNotFound)
			err := uc.ChangePassword(ctx, email, &changePassModel)
			So(err, ShouldNotBeNil)
		})
		Convey("Check email is exist", func() {
			Convey("Check old password is not correct", func() {
			})
			Convey("Check old password is correct", func() {
				mockAccountStorage.EXPECT().GetAccountByEmail(ctx, email).Return(&entities.Account{Password: "$2a$10$A.I3t3g6ocU4jkAC1F4Ojuxu2VjJ6VJV40MTv2g2vrrahSakMfGDG"}, nil)
				changePassModel.OldPassword = "123hihi"
				Convey("Update password fail", func() {
					mockAccountStorage.EXPECT().UpdateAccountById(ctx, gomock.Any(), gomock.Any()).Return(errors.New("error"))
					err := uc.ChangePassword(ctx, email, &changePassModel)
					So(err, ShouldNotBeNil)
				})
				Convey("Update password success", func() {
					mockAccountStorage.EXPECT().UpdateAccountById(ctx, gomock.Any(), gomock.Any()).Return(nil)
					err := uc.ChangePassword(ctx, email, &changePassModel)
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
