package amenityusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListAmenityByObjectID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAmenityStorage := NewMockAmenityStorage(ctrl)

	// init usecase
	cfg := config.Config{}
	uc := NewAmenityUseCase(mockAmenityStorage, &cfg)

	objectID := gofakeit.Number(1, 100)
	objectType := gofakeit.Number(1, 2)

	Convey("Test List Amenity By Object ID", t, func() {
		Convey("List Amenity By Object ID Fail", func() {
			mockAmenityStorage.EXPECT().ListByObjectID(ctx, objectID, objectType).Return(nil, errors.New("error")).Times(1)
			_, err := uc.ListAmenityByObjectID(ctx, objectID, objectType)
			So(err, ShouldNotBeNil)
		})
		Convey("List Amenity By Object ID Success", func() {
			mockAmenityStorage.EXPECT().ListByObjectID(ctx, objectID, objectType).Return(nil, nil).AnyTimes()
			_, err := uc.ListAmenityByObjectID(ctx, objectID, objectType)
			So(err, ShouldBeNil)
		})
	})
}
