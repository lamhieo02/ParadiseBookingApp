package userhandler

import "github.com/gin-gonic/gin"

type userhandler struct {
}

func (hdl *userhandler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
