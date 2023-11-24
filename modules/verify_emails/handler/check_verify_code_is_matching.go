package verifyemailshanlder

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdl *verifyEmailsHandler) CheckVerifyCodeIsMatching() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		verifyCode := c.Query("secret_code")

		err := hdl.verifyEmailsUC.CheckVerifyCodeIsMatching(c.Request.Context(), email, verifyCode)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true})
	}
}
