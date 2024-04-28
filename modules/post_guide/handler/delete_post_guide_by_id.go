package postguidehandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *postGuideHandler) DeletePostGuideByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postGuideID := ctx.Param("id")
		id, err := strconv.Atoi(postGuideID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err = h.postGuideUC.DeletePostGuideByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
