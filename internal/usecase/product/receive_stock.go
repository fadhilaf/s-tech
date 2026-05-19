package usecase

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/fadhilaf/s-tech/internal/model"
	repositoryModel "github.com/fadhilaf/s-tech/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	utils "github.com/fadhilaf/s-tech/internal/utils"
)

func (usecase *productUsecaseImpl) ReceiveProductStock(req model.ReceiveProductStockRequest) model.WebServiceResponse {
	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		return utils.ToWebServiceResponse("Product ID tidak valid", http.StatusBadRequest, nil)
	}

	supplierID, err := uuid.Parse(req.SupplierID)
	if err != nil {
		return utils.ToWebServiceResponse("Supplier ID tidak valid", http.StatusBadRequest, nil)
	}

	// Verify product exists
	_, err = usecase.Store.GetProductById(context.Background(), productID)
	if err != nil {
		return utils.ToWebServiceResponse("Produk tidak ditemukan", http.StatusNotFound, nil)
	}

	// Verify supplier exists
	_, err = usecase.Store.GetSupplierById(context.Background(), supplierID)
	if err != nil {
		return utils.ToWebServiceResponse("Supplier tidak ditemukan", http.StatusNotFound, nil)
	}

	_, err = usecase.Store.InsertProductStock(context.Background(), repositoryModel.InsertProductStockParams{
		ProductID:  productID,
		SupplierID: supplierID,
		IsAdd:      true,
		Quantity:   req.Quantity,
		Price:      req.Price,
	})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal memasukkan log stok produk ke database", http.StatusInternalServerError, nil)
	}

	// Fetch updated product to return
	updatedProduct, _ := usecase.Store.GetProductById(context.Background(), productID)

	return utils.ToWebServiceResponse("Berhasil menerima stok produk", http.StatusCreated, gin.H{
		"product": updatedProduct,
	})
}
