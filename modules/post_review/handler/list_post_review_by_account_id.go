package postreviewhandler

import (
	"net/http"
	"paradise-booking/common"
	"paradise-booking/modules/post_review/convert"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *postReviewHandler) ListPostReviewByAccountID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		page, _ := strconv.Atoi(c.Query("page"))
		limit, _ := strconv.Atoi(c.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		accountID := c.Param("account_id")
		accId, err := strconv.Atoi(accountID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := hdl.postReviewUC.ListPostReviewByAccountID(c.Request.Context(), accId, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := convert.ConvertListPostReviewToModel(data, &paging)
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
