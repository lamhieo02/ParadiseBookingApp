package amenityusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"paradise-booking/modules/amenity/iomodel"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDeleteAmenityByListId(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAmenityStorage := NewMockAmenityStorage(ctrl)

	// init usecase
	cfg := config.Config{}
	uc := NewAmenityUseCase(mockAmenityStorage, &cfg)

	testCase := iomodel.DeleteAmenityReq{}
	gofakeit.Struct(&testCase)

	Convey("Test Delete Amenity By List ID", t, func() {
		Convey("Delete Amenity By List ID Fail", func() {
			mockAmenityStorage.EXPECT().DeleteByCondition(ctx, gomock.Any()).Return(errors.New("error")).Times(1)
			err := uc.DeleteAmenityByListId(ctx, &testCase)
			So(err, ShouldNotBeNil)
		})
		Convey("Delete Amenity By List ID Success", func() {
			mockAmenityStorage.EXPECT().DeleteByCondition(ctx, gomock.Any()).Return(nil).AnyTimes()
			err := uc.DeleteAmenityByListId(ctx, &testCase)
			So(err, ShouldBeNil)
		})
	})
}
