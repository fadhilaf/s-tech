package delivery

import (
	"net/http"

	"github.com/fadhilaf/s-tech/internal/utils"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (handler *productHandler) CreateProduct(ctx *gin.Context) {
	var req model.CreateProductNoFileFormRequest

	ok := utils.BindFormAndValidate(ctx, &req)
	if !ok {
		return
	}

	supplierId, err := uuid.Parse(req.SupplierID)
	if err != nil {
		return
	}

	filename := ""
	if !handler.IsStaticCloud {
		// Simpan upload file ke folder assets/images
		filename, ok = utils.SaveFileFromForm(ctx, "image", handler.AppStaticPath+"/img/")
		if !ok {
			return
		}
	}

	productNoFile := model.ProductNoFile{
		Name:        req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
		SupplierID:  supplierId,
		IsService:   req.IsService,
		Description: req.Description,
	}

	res := handler.usecase.CreateProduct(model.CreateProductRequest{
		ProductNoFile: productNoFile,
		Image:         filename,
	})

	var path string
	path = "/admin/add-product"

	if res.Status == http.StatusCreated {
		path = "/admin/dashboard"
	}

	ctx.Header("HX-Redirect", path)

	ctx.JSON(res.Status, res)
}
