package delivery

import (
	"github.com/gin-gonic/gin"
)

func (delivery *productHandler) GetSuppliers(ctx *gin.Context) {
	res := delivery.usecase.GetSuppliers()
	ctx.JSON(res.Status, res)
}
