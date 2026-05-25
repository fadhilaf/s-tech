package usecase

import (
	"context"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
)

func (usecase *productUsecaseImpl) GetAllChronology() model.WebServiceResponse {
	chronology, err := usecase.Store.GetAllChronology(context.Background())
	if err != nil {
		return utils.ToWebServiceResponse("Gagal mengambil data laporan", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil mengambil data laporan", http.StatusOK, gin.H{
		"chronology": chronology,
	})
}
