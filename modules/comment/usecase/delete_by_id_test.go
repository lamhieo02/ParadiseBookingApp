package commentusecase

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDeleteCommentByID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockCommentSto := NewMockCommentStorage(ctrl)
	// init usecase
	uc := NewCommentUseCase(mockCommentSto, nil, nil)

	id := gofakeit.Number(1, 100)

	Convey("Test Delete Comment By ID", t, func() {
		Convey("Delete comment fail", func() {
			mockCommentSto.EXPECT().DeleteByID(ctx, id).Return(errors.New("error"))
			err := uc.DeleteCommentByID(ctx, id)
			So(err, ShouldNotBeNil)
		})
		Convey("Delete comment success", func() {
			mockCommentSto.EXPECT().DeleteByID(ctx, id).Return(nil)
			err := uc.DeleteCommentByID(ctx, id)
			So(err, ShouldBeNil)
		})
	})
}
