package usecase

import (
	"context"
	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/google/uuid"
)

func (usecase *productUsecaseImpl) DeleteSupplier(id string) model.WebServiceResponse {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return utils.ToWebServiceResponse("ID supplier tidak valid", http.StatusBadRequest, nil)
	}

	err = usecase.Store.DeleteSupplier(context.Background(), parsedId)
	if err != nil {
		// Check for foreign key constraint violation or any other errors
		return utils.ToWebServiceResponse("Gagal menghapus supplier, mungkin masih terikat pada riwayat stok", http.StatusBadRequest, nil)
	}

	return utils.ToWebServiceResponse("Berhasil menghapus supplier", http.StatusOK, nil)
}
