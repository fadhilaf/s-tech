package delivery

import (
	"net/http"
	"net/url"

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
	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya HTML
	utils.SaveResponse(ctx, res.Message)

	var location url.URL
	location = url.URL{Path: "/admin"}

	if res.Status == http.StatusOK {
		utils.SaveAdminToSession(ctx, true)
		location = url.URL{Path: "/admin/dashboard"}
	}

	ctx.Redirect(http.StatusFound, location.RequestURI())
}
