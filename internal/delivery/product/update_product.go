package delivery

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/fadhilaf/s-tech/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (delivery *productHandler) UpdateProductDetails(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	var req model.UpdateProductDetailsRequest
	req.ID = id
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle optional image
	file, err := c.FormFile("image")
	if err == nil {
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension
		if err := c.SaveUploadedFile(file, "public/img/"+newFileName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save image: %s", err.Error())})
			return
		}
		req.Image = newFileName
	}

	res := delivery.usecase.UpdateProductDetails(req)
	
	if res.Status == http.StatusOK {
		c.Header("HX-Redirect", "/admin/dashboard")
	}
	
	c.JSON(res.Status, res)
}

func (delivery *productHandler) UpdateProductPrice(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	var req model.UpdateProductPriceRequest
	req.ProductID = id
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := delivery.usecase.UpdateProductPrice(req)

	if res.Status == http.StatusCreated {
		c.Header("HX-Redirect", "/admin/dashboard")
	}

	c.JSON(res.Status, res)
}

func (delivery *productHandler) GetProductChronology(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	req := model.GetProductChronologyRequest{ProductID: id}
	res := delivery.usecase.GetProductChronology(req)

	c.JSON(res.Status, res)
}

func (delivery *productHandler) GetProductPrices(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	req := model.GetProductPricesRequest{ProductID: id}
	res := delivery.usecase.GetProductPrices(req)

	c.JSON(res.Status, res)
}
