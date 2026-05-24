package route

import (
	delivery "github.com/fadhilaf/s-tech/internal/delivery/product"
	"github.com/fadhilaf/s-tech/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.RouterGroup, handler delivery.ProductDelivery) {
	router.GET("/", handler.GetProducts)
	router.GET("/:id", handler.GetProductById)

	adminGroup := router.Group("/admin", middleware.ShouldAdmin())
	adminGroup.GET("/supplier", handler.GetSuppliers)
	adminGroup.POST("/", handler.CreateProduct)
	adminGroup.POST("/supplier", handler.CreateSupplier)
	adminGroup.DELETE("/supplier/:id", handler.DeleteSupplier)
	adminGroup.POST("/receive_stock", handler.ReceiveProductStock)
}
