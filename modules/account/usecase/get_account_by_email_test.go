package accountusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	"paradise-booking/entities"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAccountByEmail(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAccountStorage := NewMockAccountStorage(ctrl)

	// init usecase
	uc := NewUserUseCase(nil, mockAccountStorage, nil, nil, nil)

	// prepare data to test
	email := gofakeit.Email()
	Convey("Test Get Account By Email", t, func() {
		Convey("Get account by email fail", func() {
			mockAccountStorage.EXPECT().GetAccountByEmail(ctx, email).Return(nil, errors.New("error"))
			_, err := uc.GetAccountByEmail(ctx, email)
			So(err, ShouldNotBeNil)
		})
		Convey("Get account by email success", func() {
			now := time.Now()
			mockAccountStorage.EXPECT().GetAccountByEmail(ctx, email).Return(&entities.Account{
				SQLModel: common.SQLModel{CreatedAt: &now, UpdatedAt: &now},
			}, nil)
			_, err := uc.GetAccountByEmail(ctx, email)
			So(err, ShouldBeNil)
		})
	})
}
