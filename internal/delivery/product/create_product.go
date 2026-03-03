package delivery

import (
	"net/http"
	"net/url"

	"github.com/fadhilaf/s-tech/internal/utils"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *productHandler) CreateProduct(ctx *gin.Context) {
	var req model.CreateProductNoFileRequest

	ok := utils.BindFormAndValidate(ctx, &req)
	if !ok {
		return
	}

	filename := ""
	if !handler.IsStaticCloud {
		// Simpan upload file ke folder assets/images
		filename, ok = utils.SaveFileFromForm(ctx, "image", handler.AppStaticPath + "/img/")
		if !ok {
			return
		}
	}

	res := handler.usecase.CreateProduct(model.CreateProductRequest{
		NotFile: req,
		Image:   filename,
	})

	// Gaya REST API
	// ctx.JSON(res.Status, res)

	// Gaya MVC
	utils.SaveResponse(ctx, res.Message)

	var location url.URL
	location = url.URL{Path: "/admin/tambah"}

	if res.Status == http.StatusCreated {
		location = url.URL{Path: "/admin/dashboard"}
	}
	ctx.Redirect(http.StatusFound, location.RequestURI())
}
