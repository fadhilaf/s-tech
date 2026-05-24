package usecase

import (
	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/repository"
)

type ProductUsecase interface {
	CreateProduct(req model.CreateProductRequest) model.WebServiceResponse
	CreateSupplier(req model.CreateSupplierRequest) model.WebServiceResponse
	ReceiveProductStock(req model.ReceiveProductStockRequest) model.WebServiceResponse
	GetProducts(keyword string) model.WebServiceResponse
	GetProductById(id string) model.WebServiceResponse
	GetSuppliers() model.WebServiceResponse
	DeleteSupplier(id string) model.WebServiceResponse
}

var _ ProductUsecase = &productUsecaseImpl{}

func NewProductUsecase(store repository.Store) ProductUsecase {
	return &productUsecaseImpl{
		Store: store,
	}
}

type productUsecaseImpl struct {
	repository.Store
}
