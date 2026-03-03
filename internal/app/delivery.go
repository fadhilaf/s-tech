package app

import (
	auth_delivery "github.com/fadhilaf/s-tech/internal/delivery/auth"
	order_delivery "github.com/fadhilaf/s-tech/internal/delivery/order"
	product_delivery "github.com/fadhilaf/s-tech/internal/delivery/product"
	user_delivery "github.com/fadhilaf/s-tech/internal/delivery/user"
	view_delivery "github.com/fadhilaf/s-tech/internal/delivery/view"
)

type deliveries struct {
	auth    auth_delivery.AuthDelivery
	user    user_delivery.UserDelivery
	product product_delivery.ProductDelivery
	order   order_delivery.OrderDelivery
	view    view_delivery.ViewDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries
	deliveries.auth = auth_delivery.NewAuthDelivery(app.usecase.auth)
	deliveries.user = user_delivery.NewUserDelivery(app.usecase.user)
	deliveries.product = product_delivery.NewProductDelivery(app.usecase.product, app.Config.AppStaticPath, app.Config.IsStaticCloud)
	deliveries.order = order_delivery.NewOrderDelivery(app.usecase.order)
	deliveries.view = view_delivery.NewViewDelivery(app.usecase.view, app.Config.AdminPhone)

	app.delivery = deliveries
}
