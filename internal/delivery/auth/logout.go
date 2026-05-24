package delivery

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/utils"

	"github.com/gin-gonic/gin"
)

func (handler *authHandler) Logout(ctx *gin.Context) {
	utils.RemoveAuthSession(ctx)

	res := utils.ToWebServiceResponse("Logout berhasil", http.StatusOK, nil)
	ctx.Header("HX-Redirect", "/")

	ctx.JSON(res.Status, res)
}
