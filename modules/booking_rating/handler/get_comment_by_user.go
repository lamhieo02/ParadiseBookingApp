package bookingratinghandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingratinghandler) GetCommentByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userID := ctx.Query("user_id")
		id, err := strconv.Atoi(userID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		objectType := ctx.Query("object_type")
		objectTypeInt, err := strconv.Atoi(objectType)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.placeRatingUC.GetCommentByUserID(ctx, id, objectTypeInt)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res})

	}
}

// trigger cicd
