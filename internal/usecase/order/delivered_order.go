package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/fadhilaf/s-tech/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
)

func (usecase *orderUsecaseImpl) DeliveredOrder(req model.UpdateOrderStatusDeliveredRequest) model.WebServiceResponse {
	order, err := usecase.Store.GetOrderById(context.Background(), req.ID)
	if err != nil {
		return utils.ToWebServiceResponse("Order tidak ditemukan", http.StatusNotFound, nil)
	}

	if order.UserID != req.UserID {
		return utils.ToWebServiceResponse("Order bukan milik user", http.StatusUnauthorized, nil)
	}

	if order.Status == model.OrderStatusPending {
		return utils.ToWebServiceResponse("Order masih belum diproses", http.StatusUnprocessableEntity, nil)
	}

	if order.Status == model.OrderStatusDelivered {
		return utils.ToWebServiceResponse("Order sudah dikirim", http.StatusUnprocessableEntity, nil)
	}

	orderNew, err := usecase.Store.UpdateOrderStatus(context.Background(),
		repositoryModel.UpdateOrderStatusParams{ID: req.ID, Status: model.OrderStatusDelivered})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal mengupdate order ke database", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil mengupdate status order ke delivered", http.StatusOK, gin.H{"order": orderNew})
}
