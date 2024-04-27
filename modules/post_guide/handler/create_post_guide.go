package postguidehandler

import (
	"net/http"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"

	"github.com/gin-gonic/gin"
)

func (h *postGuideHandler) CreatePostGuide() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var postGuideBody postguideiomodel.CreatePostGuideReq
		if err := ctx.ShouldBind(&postGuideBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if err := h.postGuideUC.CreatePostGuide(ctx, &postGuideBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
