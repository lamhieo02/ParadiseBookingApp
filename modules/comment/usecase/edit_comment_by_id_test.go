package commentusecase

import (
	"context"
	"errors"
	"paradise-booking/entities"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEditCommentByID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockCommentSto := NewMockCommentStorage(ctrl)
	// init usecase
	uc := NewCommentUseCase(mockCommentSto, nil, nil)

	id := gofakeit.Number(1, 100)
	model := entities.Comment{}
	if err := gofakeit.Struct(&model); err != nil {
		t.Error("Error to create data test")
	}

	Convey("Test Edit Comment By ID", t, func() {
		Convey("Edit comment fail", func() {
			mockCommentSto.EXPECT().UpdateByID(ctx, id, gomock.Any()).Return(errors.New("error"))
			err := uc.EditCommentByID(ctx, id, &model)
			So(err, ShouldNotBeNil)
		})
		Convey("Edit comment success", func() {
			mockCommentSto.EXPECT().UpdateByID(ctx, id, gomock.Any()).Return(nil)
			err := uc.EditCommentByID(ctx, id, &model)
			So(err, ShouldBeNil)
		})
	})
}
