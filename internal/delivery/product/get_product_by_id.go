package delivery

import (
	"github.com/gin-gonic/gin"
)

func (delivery *productHandler) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	res := delivery.usecase.GetProductById(id)
	ctx.JSON(res.Status, res)
}
