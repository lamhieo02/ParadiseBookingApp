package userhandler

import (
	"context"
	"net/http"
	"paradise-booking/modules/user/iomodel"
	jwtprovider "paradise-booking/provider/jwt"

	"github.com/gin-gonic/gin"
)

type userUseCase interface {
	RegisterAccount(context.Context, *iomodel.AccountRegister) error
	LoginAccount(context.Context, *iomodel.AccountLogin) (*jwtprovider.Token, error)
	// UpdateProfile(context.Context, *usermodel.UserUpdate, string) error
	//GetProfileOfCurrentUser
}
type userhandler struct {
	userUC userUseCase
}

func NewUserHandler(userUseCase userUseCase) *userhandler {
	return &userhandler{userUC: userUseCase}
}

func (hdl *userhandler) RegisterAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data iomodel.AccountRegister

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if err := hdl.userUC.RegisterAccount(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data.Email})
	}
}
