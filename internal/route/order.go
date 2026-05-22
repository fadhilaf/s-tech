package route

import (
	delivery "github.com/fadhilaf/s-tech/internal/delivery/order"
	"github.com/fadhilaf/s-tech/internal/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.RouterGroup, handler delivery.OrderDelivery) {
	userRoutes := router.Group("/", middleware.ShouldUser())
	userRoutes.POST("/", handler.CreateOrder)
	userRoutes.POST("/delivered", handler.DeliveredOrder)
	userRoutes.GET("/my-orders", handler.GetMyOrders)

	adminRoutes := router.Group("/", middleware.ShouldAdmin())
	adminRoutes.POST("/processing", handler.ProcessingOrder)
	adminRoutes.GET("/", handler.GetOrders)
}
