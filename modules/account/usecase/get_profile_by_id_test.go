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

func TestGetAccountByID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAccountStorage := NewMockAccountStorage(ctrl)

	// init usecase
	uc := NewUserUseCase(nil, mockAccountStorage, nil, nil, nil)

	// prepare data to test
	id := gofakeit.Number(1, 100)

	Convey("Test Get Account By ID", t, func() {
		Convey("Get account by ID fail", func() {
			mockAccountStorage.EXPECT().GetProfileByID(ctx, id).Return(nil, errors.New("error"))
			_, err := uc.GetAccountByID(ctx, id)
			So(err, ShouldNotBeNil)
		})
		Convey("Get account by ID success", func() {
			now := time.Now()
			mockAccountStorage.EXPECT().GetProfileByID(ctx, id).Return(&entities.Account{SQLModel: common.SQLModel{CreatedAt: &now, UpdatedAt: &now}}, nil)
			_, err := uc.GetAccountByID(ctx, id)
			So(err, ShouldBeNil)
		})
	})
}
