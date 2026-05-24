package delivery

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/utils"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *authHandler) AdminLogin(ctx *gin.Context) {
	var req model.LoginRequest

	ok := utils.BindFormAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := handler.usecase.AdminLogin(req)

	if res.Status == http.StatusOK {
		utils.SaveAdminToSession(ctx, true)
		ctx.Header("HX-Redirect", "/admin/dashboard")
	}

	ctx.JSON(res.Status, res)
}
