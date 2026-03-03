package route

import (
	delivery "github.com/fadhilaf/s-tech/internal/delivery/view"
	"github.com/fadhilaf/s-tech/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ViewRoutes(router *gin.RouterGroup, handler delivery.ViewDelivery) {
	router.GET("/", handler.RenderHome)
	router.GET("/register", handler.RenderRegister)
	router.GET("/login", handler.RenderLogin)
	router.GET("/admin", handler.RenderAdmin)

	// router.GET("/blabla", middleware1, middleware2) // biso ck ini jg dibuatny

	userGroup := router.Group("/", middleware.ShouldUser())
	userGroup.GET("/pesan", handler.RenderPesan)
	userGroup.GET("/pesanan", handler.RenderPesanan)

	adminGroup := router.Group("/admin", middleware.ShouldAdmin())
	adminGroup.GET("/tambah", handler.RenderTambah)
	adminGroup.GET("/dashboard", handler.RenderDashboard)
	adminGroup.GET("/pesanan", handler.RenderAdminPesanan)
}
