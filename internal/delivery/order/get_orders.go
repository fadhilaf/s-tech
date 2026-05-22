package delivery

import (
	"github.com/gin-gonic/gin"
)

func (delivery *orderHandler) GetOrders(ctx *gin.Context) {
	res := delivery.usecase.GetOrders()
	ctx.JSON(res.Status, res)
}
