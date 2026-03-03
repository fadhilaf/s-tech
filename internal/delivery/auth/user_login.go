package delivery

import (
	"net/http"
	"net/url"

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

	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya HTML
	utils.SaveResponse(ctx, res.Message)

	var location url.URL
	location = url.URL{Path: "/login"}

	if res.Status == http.StatusOK {
		//casting dari [interface{}]interface{} ke model.User
		user, ok := res.Data["user"].(model.User)
		if ok {
			utils.SaveUserToSession(ctx, user.ID)
			location = url.URL{Path: "/"}
		}
	}

	ctx.Redirect(http.StatusFound, location.RequestURI())
}
