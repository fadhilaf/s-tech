package app

import (
	"github.com/fadhilaf/s-tech/internal/route"
	"github.com/gin-gonic/gin"
)

func (app *App) viewHandler(router *gin.RouterGroup) {
	viewGroup := router.Group("/")
	route.ViewRoutes(viewGroup, app.delivery.view)

	// viewGroup := router.Group("/", middleware1, middleware2) //biso ditaro disampingny middlewareny
}

func (app *App) apiHandler(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	route.UserRoutes(userGroup, app.delivery.user)

	productGroup := router.Group("/product")
	route.ProductRoutes(productGroup, app.delivery.product)

	orderGroup := router.Group("/order")
	route.OrderRoutes(orderGroup, app.delivery.order)

	authGroup := router.Group("/auth")
	route.AuthRoutes(authGroup, app.delivery.auth)
}
