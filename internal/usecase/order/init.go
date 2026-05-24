package usecase

import (
	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/repository"
	"github.com/google/uuid"
)

type OrderUsecase interface {
	CreateOrder(model.CreateOrderRequest) model.WebServiceResponse
	ProcessingOrder(model.UpdateOrderStatusProcessingRequest) model.WebServiceResponse
	DeliveredOrder(model.UpdateOrderStatusDeliveredRequest) model.WebServiceResponse
	GetOrders() model.WebServiceResponse
	GetOrdersByUserId(userId uuid.UUID) model.WebServiceResponse
}

var _ OrderUsecase = &orderUsecaseImpl{}

func NewOrderUsecase(store repository.Store) OrderUsecase {
	return &orderUsecaseImpl{
		Store: store,
	}
}

type orderUsecaseImpl struct {
	repository.Store
}
