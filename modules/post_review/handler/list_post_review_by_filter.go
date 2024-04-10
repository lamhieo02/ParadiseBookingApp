package postreviewhandler

import (
	"net/http"
	"paradise-booking/common"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *postReviewHandler) ListPostReviewByFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		page, _ := strconv.Atoi(c.Query("page"))
		limit, _ := strconv.Atoi(c.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		var filter postreviewiomodel.Filter
		topicID := c.Query("topic_id")
		if topicID != "" {
			topic, _ := strconv.Atoi(topicID)
			filter.TopicID = topic
		}

		accountID := c.Query("account_id")
		accId, err := strconv.Atoi(accountID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := hdl.postReviewUC.ListPostReviewByFilter(c.Request.Context(), &paging, &filter, int64(accId))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
