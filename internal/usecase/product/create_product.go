package usecase

import (
	"context"

	"net/http"

	"github.com/fadhilaf/s-tech/internal/model"
	repositoryModel "github.com/fadhilaf/s-tech/internal/repository/postgres/sqlc"
	"github.com/gin-gonic/gin"

	utils "github.com/fadhilaf/s-tech/internal/utils"
)

func (usecase *productUsecaseImpl) CreateProduct(req model.CreateProductRequest) model.WebServiceResponse {
	_, err := usecase.Store.GetProductByName(context.Background(), req.ProductNoFile.Name)
	if err == nil {
		return utils.ToWebServiceResponse("Produk dengan nama yang sama sudah ada", http.StatusConflict, nil)
	}

	productId, err := usecase.Store.CreateProduct(context.Background(), repositoryModel.CreateProductParams{
		Name:        req.ProductNoFile.Name,
		Stock:       0, // Initialize with 0, DB trigger will update it
		IsService:   req.ProductNoFile.IsService,
		Description: req.ProductNoFile.Description,
		Image:       req.Image,
	})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal memasukkan produk ke database", http.StatusInternalServerError, nil)
	}

	_, err = usecase.Store.UpdateProductPrice(context.Background(), repositoryModel.UpdateProductPriceParams{
		ProductID: productId,
		Price:     req.ProductNoFile.Price,
	})
	if err != nil {
		return utils.ToWebServiceResponse("Gagal memasukkan harga produk ke database", http.StatusInternalServerError, nil)
	}

	// Insert initial stock using the provided SupplierID
	if !req.ProductNoFile.IsService {
		_, err = usecase.Store.InsertProductStock(context.Background(), repositoryModel.InsertProductStockParams{
			ProductID:  productId,
			SupplierID: req.ProductNoFile.SupplierID,
			IsAdd:      true,
			Quantity:   req.ProductNoFile.Stock,
			Price:      req.ProductNoFile.Price, // Cost price can be same as selling price for now or 0
		})
		if err != nil {
			return utils.ToWebServiceResponse("Gagal menambahkan stok awal produk", http.StatusInternalServerError, nil)
		}
	}

	return utils.ToWebServiceResponse("Berhasil memasukkan produk ke database", http.StatusCreated, gin.H{
		"product_id": productId,
	})
}
