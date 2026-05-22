package delivery

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/utils"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *authHandler) UserLogin(ctx *gin.Context) {
	var req model.LoginRequest

	ok := utils.BindFormAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := handler.usecase.UserLogin(req)

	if res.Status == http.StatusNotFound || res.Status == http.StatusUnauthorized {
		if errors, ok := res.Data["errors"].(map[string]string); ok {
			ctx.JSON(res.Status, utils.ToDetailedErrorWebServiceResponse(res.Message, res.Status, errors))
			return
		}
	}

	if res.Status == http.StatusOK {
		//casting dari [interface{}]interface{} ke model.User
		if user, ok := res.Data["user"].(model.User); ok {
			utils.SaveUserToSession(ctx, user.ID)
		}

	}

	ctx.JSON(res.Status, res)
}
