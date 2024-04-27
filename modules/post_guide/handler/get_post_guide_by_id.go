package postguidehandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *postGuideHandler) GetPostGuideByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postGuideID := ctx.Param("id")
		id, err := strconv.Atoi(postGuideID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		data, err := h.postGuideUC.GetPostGuideByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
