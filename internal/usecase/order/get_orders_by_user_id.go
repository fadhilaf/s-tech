package usecase

import (
	"context"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (usecase *orderUsecaseImpl) GetOrdersByUserId(userId uuid.UUID) model.WebServiceResponse {
	ordersDb, err := usecase.Store.GetOrdersByUserId(context.Background(), userId)
	if err != nil {
		return utils.ToWebServiceResponse("Orders tidak ditemukan", http.StatusNotFound, nil)
	}

	orders := make([]model.Order, len(ordersDb))

	for i, order := range ordersDb {
		product, err := usecase.Store.GetProductByPriceId(context.Background(), order.ProductPriceID)
		if err != nil {
			return utils.ToWebServiceResponse("Product tidak ditemukan di database", http.StatusNotFound, nil)
		}

		orders[i] = model.Order{
			ID:             order.ID,
			ProductPriceID: order.ProductPriceID,
			ProductID:      product.ID,
			ProductName:    product.Name,
			ProductPrice:   product.CurrentPrice,
			IsService:      product.IsService,
			Quantity:       order.Quantity,
			Status:         string(order.Status.([]uint8)),
			Description:    order.Description,
		}
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan orders", http.StatusOK, gin.H{
		"orders": orders,
	})
}
