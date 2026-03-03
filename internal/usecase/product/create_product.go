package usecase

import (
	"context"

	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	repositoryModel "github.com/fadhilaf/s-tech/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	utils "github.com/fadhilaf/s-tech/internal/utils"
)

func (usecase *productUsecaseImpl) CreateProduct(req model.CreateProductRequest) model.WebServiceResponse {
	_, err := usecase.Store.GetProductByName(context.Background(), req.NotFile.Name)
	if err == nil {
		return utils.ToWebServiceResponse("Produk dengan nama yang sama sudah ada", http.StatusConflict, nil)
	}

	product, err := usecase.Store.CreateProduct(context.Background(), repositoryModel.CreateProductParams{
		Name:        req.NotFile.Name,
		Price:       req.NotFile.Price,
		Stock:       req.NotFile.Stock,
		IsService:   req.NotFile.IsService,
		Description: req.NotFile.Description,
		Image:       req.Image,
	})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal memasukkan produk ke database", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil memasukkan produk ke database", http.StatusCreated, gin.H{
		"product": product,
	})
}
