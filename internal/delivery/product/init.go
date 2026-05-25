package delivery

import (
	usecase "github.com/fadhilaf/s-tech/internal/usecase/product"
	"github.com/gin-gonic/gin"
)

type ProductDelivery interface {
	CreateProduct(ctx *gin.Context)
	CreateSupplier(ctx *gin.Context)
	ReceiveProductStock(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	GetSuppliers(ctx *gin.Context)
	DeleteSupplier(ctx *gin.Context)
	UpdateProductDetails(ctx *gin.Context)
	UpdateProductPrice(ctx *gin.Context)
	GetProductPrices(ctx *gin.Context)
	GetProductChronology(ctx *gin.Context)
	GetAllChronology(ctx *gin.Context)
}

var _ ProductDelivery = &productHandler{}

func NewProductDelivery(usecase usecase.ProductUsecase, appStaticPath string, isStaticCloud bool) ProductDelivery {
	return &productHandler{
		usecase: usecase,
		AppStaticPath: appStaticPath,
		IsStaticCloud: isStaticCloud,
	}
}

// type userHandler ini ditambahi satu per satu per file selain init
type productHandler struct {
	usecase usecase.ProductUsecase
	AppStaticPath string
	IsStaticCloud bool
}

