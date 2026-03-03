package usecase

import (
	"context"
	"net/http"

	repositoryModel "github.com/fadhilaf/s-tech/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
)

func (usecase *orderUsecaseImpl) ProcessingOrder(req model.UpdateOrderStatusProcessingRequest) model.WebServiceResponse {
	order, err := usecase.Store.GetOrderById(context.Background(), req.ID)
	if err != nil {
		return utils.ToWebServiceResponse("Order tidak ditemukan", http.StatusNotFound, nil)
	}

	if order.Status == model.OrderStatusProcessing {
		return utils.ToWebServiceResponse("Order sudah diproses", http.StatusUnprocessableEntity, nil)
	}

	if order.Status == model.OrderStatusDelivered {
		return utils.ToWebServiceResponse("Order sudah dikirim", http.StatusUnprocessableEntity, nil)
	}

	orderNew, err := usecase.Store.UpdateOrderStatus(context.Background(), repositoryModel.UpdateOrderStatusParams{ID: req.ID, Status: model.OrderStatusProcessing})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal mengupdate order ke database", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil mengupdate status order ke processing", http.StatusOK, gin.H{"order": orderNew})
}
