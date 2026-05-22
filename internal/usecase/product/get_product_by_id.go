package usecase

import (
	"context"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (usecase *productUsecaseImpl) GetProductById(id string) model.WebServiceResponse {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return utils.ToWebServiceResponse("ID Produk tidak valid", http.StatusBadRequest, nil)
	}

	productDb, err := usecase.Store.GetProductById(context.Background(), parsedId)
	if err != nil {
		return utils.ToWebServiceResponse("Produk tidak ditemukan", http.StatusNotFound, nil)
	}

	product := model.Product{
		ID:           productDb.ID,
		Name:         productDb.Name,
		CurrentPrice: productDb.CurrentPrice,
		IsService:    productDb.IsService,
		Image:        productDb.Image,
		Description:  productDb.Description,
		Stock:        productDb.Stock,
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan detail produk", http.StatusOK, gin.H{
		"product": product,
	})
}
