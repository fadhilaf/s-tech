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
	adminGroup.PUT("/:id", handler.UpdateProductDetails)
	adminGroup.POST("/:id/price", handler.UpdateProductPrice)
	adminGroup.GET("/:id/prices", handler.GetProductPrices)
	adminGroup.GET("/:id/chronology", handler.GetProductChronology)
	adminGroup.GET("/report", handler.GetAllChronology)
	adminGroup.POST("/supplier", handler.CreateSupplier)
	adminGroup.DELETE("/supplier/:id", handler.DeleteSupplier)
	adminGroup.POST("/receive_stock", handler.ReceiveProductStock)
}
