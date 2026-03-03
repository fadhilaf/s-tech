package middleware

import (
	"github.com/fadhilaf/s-tech/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveAndLoadSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := utils.GetUserIdFromSession(c)

		if userId != uuid.Nil {
			c.Set("user_id", userId)
		}
		c.Set("is_admin", utils.GetAdminFromSession(c))

	}
}

func ShouldUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := utils.GetUserIdFromContext(c)

		if userId == uuid.Nil {
			c.HTML(401, "401.gohtml", gin.H{})
			c.Abort()
			return
		}
	}
}

func ShouldAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin := utils.GetAdminFromContext(c)

		if !isAdmin {
			c.AbortWithStatusJSON(401, gin.H{"error": "Kamu harus admin untuk mengakses halaman ini"})
			return
		}
	}
}
