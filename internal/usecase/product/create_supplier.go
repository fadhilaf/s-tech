package usecase

import (
	"context"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/gin-gonic/gin"

	utils "github.com/fadhilaf/s-tech/internal/utils"
)

func (usecase *productUsecaseImpl) CreateSupplier(req model.CreateSupplierRequest) model.WebServiceResponse {
	supplierId, err := usecase.Store.CreateSupplier(context.Background(), req.Name)
	if err != nil {
		return utils.ToWebServiceResponse("Gagal menambahkan supplier ke database", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil menambahkan supplier", http.StatusCreated, gin.H{
		"supplier_id": supplierId,
		"name":        req.Name,
	})
}
