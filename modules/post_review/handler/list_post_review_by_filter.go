package postreviewhandler

import (
	"net/http"
	"paradise-booking/common"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (hdl *postReviewHandler) ListPostReviewByFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		page, _ := strconv.Atoi(c.Query("page"))
		limit, _ := strconv.Atoi(c.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		// HANDLE FILTER
		var filter postreviewiomodel.Filter
		topicID := c.Query("topic_id")
		if topicID != "" {
			topic, _ := strconv.Atoi(topicID)
			filter.TopicID = topic
		}

		lat := c.Query("lat")
		if lat != "" {
			latFloat, _ := strconv.ParseFloat(lat, 64)
			filter.Lat = latFloat
		}

		lng := c.Query("lng")
		if lng != "" {
			lngFloat, _ := strconv.ParseFloat(lng, 64)
			filter.Lng = lngFloat
		}

		dateFrom := c.Query("date_from")
		if dateFrom != "" {
			dateFrom, err := time.Parse("2006-01-02", dateFrom)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			filter.DateFrom = &dateFrom
		}

		dateTo := c.Query("date_to")
		if dateTo != "" {
			dateTo, err := time.Parse("2006-01-02", dateTo)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			filter.DateTo = &dateTo
		}

		// END HANDLE FILTER

		result, err := hdl.postReviewUC.ListPostReviewByFilter(c.Request.Context(), &paging, &filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
