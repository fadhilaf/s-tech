package route

import (
	delivery "github.com/fadhilaf/s-tech/internal/delivery/user"
	"github.com/fadhilaf/s-tech/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, handler delivery.UserDelivery) {
	router.POST("/", handler.CreateUser)
	
	userRoutes := router.Group("/", middleware.ShouldUser())
	userRoutes.GET("/me", handler.GetProfile)
	
	// router.GET("/", handler.GetAllUser)
	// router.PUT("/:id", handler.UpdateUser)
	// router.DELETE("/:id", handler.DeleteUser)
}
