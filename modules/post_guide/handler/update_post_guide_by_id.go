package postguidehandler

import (
	"net/http"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *postGuideHandler) UpdatePostGuideByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postGuideID := ctx.Query("id")
		id, err := strconv.Atoi(postGuideID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		var postGuideBody postguideiomodel.UpdatePostGuideReq
		if err := ctx.ShouldBind(&postGuideBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err = h.postGuideUC.UpdatePostGuideByID(ctx, id, &postGuideBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
