package usecase

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fadhilaf/s-tech/internal/model"

	utils "github.com/fadhilaf/s-tech/internal/utils"
)

func (usecase *viewUsecaseImpl) GetProductById(req model.GetProductByIdRequest) model.WebServiceResponse {
	productDb, err := usecase.Store.GetProductById(context.Background(), req.ID)
	if err != nil {
		return utils.ToWebServiceResponse("Product tidak ditemukan", http.StatusNotFound, nil)
	}

	product := model.Product{
		ID:          productDb.ID,
		Name:        productDb.Name,
		Price:       utils.AddCommas(int(productDb.Price)),
		IsService:   productDb.IsService,
		Image:       productDb.Image,
		Description: productDb.Description,
		Stock:       productDb.Stock,
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan product", http.StatusOK, gin.H{
		"product": product,
	})
}
