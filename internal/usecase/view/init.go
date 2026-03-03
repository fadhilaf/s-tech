package usecase

import (
	"github.com/fadhilaf/s-tech/internal/repository"

	"github.com/fadhilaf/s-tech/internal/model"
)

type ViewUsecase interface {
	GetUserById(model.GetUserByIdRequest) model.WebServiceResponse
	GetOrder() model.WebServiceResponse
	GetOrderByUserId(model.GetOrderByUserIdRequest) model.WebServiceResponse
	GetProduct() model.WebServiceResponse
	GetProductById(model.GetProductByIdRequest) model.WebServiceResponse
	GetProductByKeyword(model.GetProductByKeywordRequest) model.WebServiceResponse
}

var _ ViewUsecase = &viewUsecaseImpl{}

func NewViewUsecase(store repository.Store) ViewUsecase {
	return &viewUsecaseImpl{
		Store: store,
	}
}

type viewUsecaseImpl struct {
	repository.Store
}
