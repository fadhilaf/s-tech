package usecase

import (
	"context"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
)

func (usecase *productUsecaseImpl) GetProducts(keyword string) model.WebServiceResponse {
	var products []model.Product

	if keyword != "" {
		productsDb, err := usecase.Store.GetProductByQuery(context.Background(), "%"+keyword+"%")
		if err != nil {
			return utils.ToWebServiceResponse("Product tidak ditemukan", http.StatusNotFound, nil)
		}
		products = make([]model.Product, len(productsDb))
		for i, product := range productsDb {
			products[i] = model.Product{
				ID:           product.ID,
				Name:         product.Name,
				CurrentPrice: product.CurrentPrice,
				Stock:        product.Stock,
				IsService:    product.IsService,
				Description:  product.Description,
				Image:        product.Image,
			}
		}
	} else {
		productsDb, err := usecase.Store.GetProduct(context.Background())
		if err != nil {
			return utils.ToWebServiceResponse("Product tidak ditemukan", http.StatusNotFound, nil)
		}
		products = make([]model.Product, len(productsDb))
		for i, product := range productsDb {
			products[i] = model.Product{
				ID:           product.ID,
				Name:         product.Name,
				CurrentPrice: product.CurrentPrice,
				Stock:        product.Stock,
				IsService:    product.IsService,
				Description:  product.Description,
				Image:        product.Image,
			}
		}
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan products", http.StatusOK, gin.H{
		"products": products,
	})
}
