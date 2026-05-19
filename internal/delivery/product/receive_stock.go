package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
)

func (delivery *productHandler) ReceiveProductStock(ctx *gin.Context) {
	var req model.ReceiveProductStockRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res := delivery.usecase.ReceiveProductStock(req)
	ctx.JSON(res.Status, res)
}
