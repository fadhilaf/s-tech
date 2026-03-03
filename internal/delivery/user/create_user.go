package delivery

import (
	"net/http"
	"net/url"

	"github.com/fadhilaf/s-tech/internal/utils"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *userHandler) CreateUser(ctx *gin.Context) {
	var req model.CreateUserRequest

	ok := utils.BindFormAndValidate(ctx, &req)
	if !ok {
		return
	}
	res := handler.usecase.CreateUser(req)

	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya MVC
	utils.SaveResponse(ctx, res.Message)

	var location url.URL
	location = url.URL{Path: "/register"}

	if res.Status == http.StatusCreated {
		location = url.URL{Path: "/login"}
	}

	ctx.Redirect(http.StatusFound, location.RequestURI())
}
