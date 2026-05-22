package delivery

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
)

func (delivery *userHandler) GetProfile(ctx *gin.Context) {
	userId, exists := ctx.Get("UserId")
	if !exists {
		res := utils.ToWebServiceResponse("Gagal mendapatkan User ID dari konteks", http.StatusUnauthorized, nil)
		ctx.JSON(res.Status, res)
		return
	}

	userIdString := userId.(string)
	res := delivery.usecase.GetProfile(userIdString)
	ctx.JSON(res.Status, res)
}
