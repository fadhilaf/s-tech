package app

import (
	"github.com/fadhilaf/s-tech/internal/repository"

	auth_usecase "github.com/fadhilaf/s-tech/internal/usecase/auth"
	order_usecase "github.com/fadhilaf/s-tech/internal/usecase/order"
	product_usecase "github.com/fadhilaf/s-tech/internal/usecase/product"
	user_usecase "github.com/fadhilaf/s-tech/internal/usecase/user"
	view_usecase "github.com/fadhilaf/s-tech/internal/usecase/view"
)

type usecases struct {
	auth    auth_usecase.AuthUsecase
	user    user_usecase.UserUsecase
	product product_usecase.ProductUsecase
	order   order_usecase.OrderUsecase
	view    view_usecase.ViewUsecase
}

func (app *App) initUsecase() {
	store := repository.NewPostgresStore(app.db)

	var usecases usecases
	usecases.auth = auth_usecase.NewAuthUsecase(store, app.Config.AdminEmail, app.Config.AdminPassword)
	usecases.user = user_usecase.NewUserUsecase(store)
	usecases.product = product_usecase.NewProductUsecase(store)
	usecases.order = order_usecase.NewOrderUsecase(store)
	usecases.view = view_usecase.NewViewUsecase(store)

	app.usecase = usecases
}
