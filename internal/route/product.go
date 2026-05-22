package route

import (
	delivery "github.com/fadhilaf/s-tech/internal/delivery/product"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.RouterGroup, handler delivery.ProductDelivery) {
	router.POST("/", handler.CreateProduct)
	router.POST("/supplier", handler.CreateSupplier)
	router.POST("/receive_stock", handler.ReceiveProductStock)
	
	router.GET("/", handler.GetProducts)
	router.GET("/supplier", handler.GetSuppliers)
	router.GET("/:id", handler.GetProductById)
}
