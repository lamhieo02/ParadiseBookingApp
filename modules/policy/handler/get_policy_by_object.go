package policieshandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *policyHandler) GetPolicyByObjectId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		objectId := ctx.Query("object_id")
		objectIdInt, _ := strconv.Atoi(objectId)

		objectType := ctx.Query("object_type")
		objectTypeInt, _ := strconv.Atoi(objectType)

		res, err := hdl.policyUC.GetPolicyByObjectID(ctx, objectIdInt, objectTypeInt)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res})
	}
}
