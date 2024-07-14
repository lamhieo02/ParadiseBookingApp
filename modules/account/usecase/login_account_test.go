package accountusecase

import (
	"context"
	"paradise-booking/config"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

func TestLoginAccount(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAccountStorage := NewMockAccountStorage(ctrl)

	// init usecase
	cfg := config.Config{}
	cfg.App.Secret = "123456"
	uc := NewUserUseCase(&cfg, mockAccountStorage, nil, nil, nil)

	// prepare data to test
	testcases := make([]*iomodel.AccountLogin, 1)
	for i := 0; i < 1; i++ {
		testcases[i] = &iomodel.AccountLogin{
			Email:    gofakeit.Email(),
			Password: gofakeit.Password(true, true, true, true, false, 10),
			Type:     gofakeit.Number(0, 1),
			FullName: gofakeit.Name(),
			Avatar:   gofakeit.URL(),
		}
	}

	for _, tc := range testcases {
		Convey("Test Login Account", t, func() {
			Convey("Check email is not exist", func() {
				mockAccountStorage.EXPECT().GetAccountByEmail(ctx, tc.Email).Return(&entities.Account{}, gorm.ErrRecordNotFound)
				result, err := uc.LoginAccount(ctx, tc)
				So(result, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("Check email is exist", func() {
				Convey("Check password is not correct", func() {
					mockAccountStorage.EXPECT().GetAccountByEmail(ctx, tc.Email).Return(&entities.Account{Email: tc.Email, Password: "$2a$10$A.I3t3g6ocU4jkAC1F4Ojuxu2VjJ6VJV40MTv2g2vrrahSakMfGDG"}, nil)
					tc.Password = "123hihi"
					result, err := uc.LoginAccount(ctx, tc)
					So(result, ShouldBeNil)
					So(err, ShouldNotBeNil)
				})
				Convey("Check password is correct", func() {
					tc.Password = "123hihi"
					Convey("Login fail", func() {
						mockAccountStorage.EXPECT().GetAccountByEmail(ctx, tc.Email).Return(&entities.Account{Email: tc.Email, Password: "$2a$10$A.I3t3g6ocU4jkAC1F4Ojuxu2VjJ6VJV40MTv2g2vrrahSakMfGDG", Status: constant.StatusInactive}, nil)
						result, err := uc.LoginAccount(ctx, tc)
						So(result, ShouldBeNil)
						So(err, ShouldNotBeNil)
					})
					Convey("Login success", func() {
						mockAccountStorage.EXPECT().GetAccountByEmail(ctx, tc.Email).Return(&entities.Account{Email: tc.Email, Password: "$2a$10$A.I3t3g6ocU4jkAC1F4Ojuxu2VjJ6VJV40MTv2g2vrrahSakMfGDG", Status: constant.StatusActive, IsEmailVerified: constant.StatusActive}, nil)
						result, err := uc.LoginAccount(ctx, tc)
						So(result, ShouldNotBeNil)
						So(err, ShouldBeNil)
					})
				})
			})
		})
	}
}

func TestLoginAccountGoogle(t *testing.T) {
	// skip this test
	t.Skip()
}
