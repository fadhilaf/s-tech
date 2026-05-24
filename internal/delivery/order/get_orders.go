package delivery

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (delivery *orderHandler) GetOrders(ctx *gin.Context) {

	if isAdmin := utils.GetAdminFromContext(ctx); isAdmin {
		res := delivery.usecase.GetOrders()
		ctx.JSON(res.Status, res)
		return
	}

	userId := utils.GetUserIdFromContext(ctx)
	if userId == uuid.Nil {
		res := utils.ToWebServiceResponse("Gagal mendapatkan User ID dari konteks", http.StatusUnauthorized, nil)
		ctx.JSON(res.Status, res)
		return
	}

	res := delivery.usecase.GetOrdersByUserId(userId)
	ctx.JSON(res.Status, res)
}
