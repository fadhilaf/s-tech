package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/fadhilaf/s-tech/internal/model"
	repositoryModel "github.com/fadhilaf/s-tech/internal/repository/postgres/sqlc"
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
)

func (usecase *productUsecaseImpl) UpdateProductDetails(req model.UpdateProductDetailsRequest) model.WebServiceResponse {
	// If image is empty, we should get the current product to retain its image
	var image string = req.Image
	if image == "" {
		existing, err := usecase.Store.GetProductById(context.Background(), req.ID)
		if err != nil {
			return utils.ToWebServiceResponse("Produk tidak ditemukan", http.StatusNotFound, nil)
		}
		image = existing.Image
	}

	_, err := usecase.Store.UpdateProductDetails(context.Background(), repositoryModel.UpdateProductDetailsParams{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Image:       image,
	})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal mengubah data produk", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil mengubah data produk", http.StatusOK, nil)
}

func (usecase *productUsecaseImpl) UpdateProductPrice(req model.UpdateProductPriceRequest) model.WebServiceResponse {
	// Parse effective date
	effectiveDate, err := time.Parse("2006-01-02T15:04", req.EffectiveDate)
	if err != nil {
		// Fallback to Date only if they didn't submit time
		effectiveDate, err = time.Parse("2006-01-02", req.EffectiveDate)
		if err != nil {
			return utils.ToWebServiceResponse("Format tanggal tidak valid", http.StatusBadRequest, nil)
		}
	}

	_, err = usecase.Store.CreateProductPrice(context.Background(), repositoryModel.CreateProductPriceParams{
		ProductID:    req.ProductID,
		Price:        req.Price,
		EffectiveDate: effectiveDate,
	})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal menambah harga produk", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil menjadwalkan harga produk", http.StatusCreated, nil)
}

func (usecase *productUsecaseImpl) GetProductChronology(req model.GetProductChronologyRequest) model.WebServiceResponse {
	chronology, err := usecase.Store.GetProductChronology(context.Background(), req.ProductID)
	if err != nil {
		return utils.ToWebServiceResponse("Gagal mengambil data kronologi", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil mengambil data kronologi", http.StatusOK, gin.H{
		"chronology": chronology,
	})
}

func (usecase *productUsecaseImpl) GetProductPrices(req model.GetProductPricesRequest) model.WebServiceResponse {
	prices, err := usecase.Store.GetProductPricesByProductId(context.Background(), req.ProductID)
	if err != nil {
		return utils.ToWebServiceResponse("Gagal mengambil riwayat harga", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Berhasil mengambil riwayat harga", http.StatusOK, gin.H{
		"prices": prices,
	})
}
