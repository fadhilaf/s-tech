package usecase

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fadhilaf/s-tech/internal/model"

	utils "github.com/fadhilaf/s-tech/internal/utils"
)

func (usecase *viewUsecaseImpl) GetProduct() model.WebServiceResponse {
	productsDb, err := usecase.Store.GetProduct(context.Background())
	if err != nil {
		return utils.ToWebServiceResponse("Product tidak ditemukan", http.StatusNotFound, nil)
	}

	products := make([]model.Product, len(productsDb))

	for i, product := range productsDb {
		products[i] = model.Product{
			ID:          product.ID,
			Name:        product.Name,
			CurrentPrice: product.CurrentPrice,
			Stock:       product.Stock,
			IsService:   product.IsService,
			Description: product.Description,
			Image:       product.Image,
		}
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan products", http.StatusOK, gin.H{
		"products": products,
	})
}
