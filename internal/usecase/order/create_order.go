package usecase

import (
	"context"

	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	repositoryModel "github.com/fadhilaf/s-tech/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	utils "github.com/fadhilaf/s-tech/internal/utils"
)

func (usecase *orderUsecaseImpl) CreateOrder(req model.CreateOrderRequest) model.WebServiceResponse {
	product, err := usecase.Store.GetProductById(context.Background(), req.ProductID)
	if err != nil {
		return utils.ToWebServiceResponse("Produk tidak ditemukan", http.StatusNotFound, nil)
	}

	if product.Stock < req.Quantity {
		return utils.ToWebServiceResponse("Stok produk tidak mencukupi", http.StatusBadRequest, nil)
	}

	order, err := usecase.Store.CreateOrder(context.Background(), repositoryModel.CreateOrderParams{
		UserID:         req.UserID,
		ProductPriceID: product.ProductPriceID,
		Quantity:       req.Quantity,
		Description:    req.Description,
	})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal memasukkan order ke database", http.StatusInternalServerError, nil)
	}

	_, err = usecase.Store.InsertProductStock(context.Background(), repositoryModel.InsertProductStockParams{
		ProductID:  req.ProductID,
		SupplierID: uuid.Nil,
		IsAdd:      false,
		Quantity:   req.Quantity,
		Price:      0, // the cost price doesn't matter for removal
	})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal memasukkan log stok produk ke database", http.StatusInternalServerError, nil)
	}

    // Refresh the product to get the updated stock
	updatedProduct, _ := usecase.Store.GetProductById(context.Background(), req.ProductID)

	return utils.ToWebServiceResponse("Berhasil memasukkan order ke database", http.StatusCreated, gin.H{
		"order":   order,
		"product": updatedProduct,
	})
}
