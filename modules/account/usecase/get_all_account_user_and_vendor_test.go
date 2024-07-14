package accountusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAllAccountUserAndVendor(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAccountStorage := NewMockAccountStorage(ctrl)

	// init usecase
	uc := NewUserUseCase(nil, mockAccountStorage, nil, nil, nil)

	// prepare data to test
	paging := common.Paging{
		Page:  gofakeit.Number(1, 10),
		Limit: gofakeit.Number(1, 10),
	}

	Convey("Test Get All Account User And Vendor", t, func() {
		Convey("Get all account user and vendor fail", func() {
			mockAccountStorage.EXPECT().GetAllAccountUserAndVendor(ctx, gomock.Any()).Return(nil, errors.New("error"))
			_, err := uc.GetAllAccountUserAndVendor(ctx, &paging)
			So(err, ShouldNotBeNil)
		})
		Convey("Get all account user and vendor success", func() {
			mockAccountStorage.EXPECT().GetAllAccountUserAndVendor(ctx, gomock.Any()).Return(nil, nil)
			_, err := uc.GetAllAccountUserAndVendor(ctx, &paging)
			So(err, ShouldBeNil)
		})
	})
}
