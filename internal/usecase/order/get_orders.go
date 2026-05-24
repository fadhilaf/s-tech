package usecase

import (
	"context"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
)

func (usecase *orderUsecaseImpl) GetOrders() model.WebServiceResponse {
	ordersDb, err := usecase.Store.GetOrders(context.Background())
	if err != nil {
		return utils.ToWebServiceResponse("Orders tidak ditemukan", http.StatusNotFound, nil)
	}

	orders := make([]model.Order, len(ordersDb))

	for i, order := range ordersDb {
		product, err := usecase.Store.GetProductByPriceId(context.Background(), order.ProductPriceID)
		if err != nil {
			return utils.ToWebServiceResponse("Product tidak ditemukan di database", http.StatusNotFound, nil)
		}

		user, err := usecase.Store.GetUserById(context.Background(), order.UserID)
		if err != nil {
			return utils.ToWebServiceResponse("User tidak ditemukan di database", http.StatusNotFound, nil)
		}

		orders[i] = model.Order{
			ID:             order.ID,
			ProductPriceID: order.ProductPriceID,
			ProductID:      product.ID,
			ProductName:    product.Name,
			ProductPrice:   product.CurrentPrice,
			IsService:      product.IsService,
			BuyerID:        user.ID,
			BuyerName:      user.Name,
			BuyerAddress:   user.Address,
			BuyerPhone:     user.Phone,
			Quantity:       order.Quantity,
			Status:         string(order.Status.([]uint8)),
			Description:    order.Description,
		}
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan orders", http.StatusOK, gin.H{
		"orders": orders,
	})
}
