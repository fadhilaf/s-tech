package delivery

import (
	"github.com/gin-gonic/gin"
)

func (delivery *productHandler) GetAllChronology(c *gin.Context) {
	res := delivery.usecase.GetAllChronology()
	c.JSON(res.Status, res)
}
