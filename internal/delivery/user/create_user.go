package delivery

import (
	"net/http"

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

	if res.Status == http.StatusConflict {
		if errors, ok := res.Data["errors"].(map[string]string); ok {
			detailedRes := utils.ToDetailedErrorWebServiceResponse(res.Message, http.StatusConflict, errors)
			ctx.JSON(detailedRes.Status, detailedRes)
			return
		}

	} else if res.Status == http.StatusCreated || res.Status == http.StatusOK {
		ctx.Header("HX-Redirect", "/login")
	}

	ctx.JSON(res.Status, res)
}
