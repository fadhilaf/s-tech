package delivery

import (
	"github.com/gin-gonic/gin"
)

func (delivery *productHandler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")
	res := delivery.usecase.DeleteSupplier(id)
	c.JSON(res.Status, res)
}
