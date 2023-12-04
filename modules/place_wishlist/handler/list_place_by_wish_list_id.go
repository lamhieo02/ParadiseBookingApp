package placewishlisthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeWishListHandler) ListPlaceByWishListID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_wishListID := ctx.Query("wish_list_id")
		wishListID, err := strconv.Atoi(_wishListID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.placeWishListUC.GetPlaceByWishListID(ctx.Request.Context(), wishListID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": res})

	}
}
