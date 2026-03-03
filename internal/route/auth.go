package route

import (
	delivery "github.com/fadhilaf/s-tech/internal/delivery/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, handler delivery.AuthDelivery) {
	router.POST("/login", handler.UserLogin)
	router.POST("/admin", handler.AdminLogin)
	router.POST("/logout", handler.Logout)
}
