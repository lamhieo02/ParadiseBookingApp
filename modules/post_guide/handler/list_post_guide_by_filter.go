package postguidehandler

import (
	"net/http"
	"paradise-booking/common"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *postGuideHandler) ListPostGuideByFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter postguideiomodel.Filter

		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		// HANDLE FILTER
		postOwnerID := ctx.Query("post_owner_id")
		if postOwnerID != "" {
			postOwner, _ := strconv.Atoi(postOwnerID)
			filter.PostOwnerId = postOwner
		}

		topicID := ctx.Query("topic_id")
		if topicID != "" {
			topic, _ := strconv.Atoi(topicID)
			filter.TopicID = topic
		}

		lat := ctx.Query("lat")
		if lat != "" {
			latFloat, _ := strconv.ParseFloat(lat, 64)
			filter.Lat = latFloat
		}

		lng := ctx.Query("lng")
		if lng != "" {
			lngFloat, _ := strconv.ParseFloat(lng, 64)
			filter.Lng = lngFloat
		}

		state := ctx.Query("state")
		if state != "" {
			filter.State = state
		}

		data, err := h.postGuideUC.ListPostGuideByFilter(ctx, &paging, &filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
