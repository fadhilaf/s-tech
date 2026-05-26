package delivery

import (
	"github.com/gin-gonic/gin"
)

func (delivery *productHandler) GetSuppliers(ctx *gin.Context) {
	res := delivery.usecase.GetSuppliers()
	ctx.JSON(res.Status, res)

	//For testing
	// ctx.JSON(http.StatusServiceUnavailable, model.WebServiceResponse{
	// 	Message: "API is turned off for testing",
	// 	Status:  http.StatusServiceUnavailable,
	// 	Data:    nil,
	// })
}
