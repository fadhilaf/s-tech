package delivery

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
)

func (delivery *userHandler) GetProfile(ctx *gin.Context) {
	userId := utils.GetUserIdFromContext(ctx)
	if userId == uuid.Nil {
		res := utils.ToWebServiceResponse("Gagal mendapatkan User ID dari konteks", http.StatusUnauthorized, nil)
		ctx.JSON(res.Status, res)
		return
	}

	res := delivery.usecase.GetProfile(userId)
	ctx.JSON(res.Status, res)
}
