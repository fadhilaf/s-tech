package usecase

import (
	"context"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
)

func (usecase *productUsecaseImpl) GetSuppliers() model.WebServiceResponse {
	suppliersDb, err := usecase.Store.GetSuppliers(context.Background())
	if err != nil {
		return utils.ToWebServiceResponse("Gagal mengambil data supplier", http.StatusInternalServerError, nil)
	}

	suppliers := make([]model.Supplier, len(suppliersDb))
	for i, s := range suppliersDb {
		suppliers[i] = model.Supplier{
			ID:   s.ID,
			Name: s.Name,
		}
	}

	return utils.ToWebServiceResponse("Berhasil mendapatkan suppliers", http.StatusOK, gin.H{
		"suppliers": suppliers,
	})
}
