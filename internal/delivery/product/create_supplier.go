package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
)

func (delivery *productHandler) CreateSupplier(ctx *gin.Context) {
	var req model.CreateSupplierRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res := delivery.usecase.CreateSupplier(req)
	ctx.JSON(res.Status, res)
}
