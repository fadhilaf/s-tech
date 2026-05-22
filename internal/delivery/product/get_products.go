package delivery

import (
	"github.com/gin-gonic/gin"
)

func (delivery *productHandler) GetProducts(ctx *gin.Context) {
	search := ctx.Query("search")
	res := delivery.usecase.GetProducts(search)
	ctx.JSON(res.Status, res)
}
