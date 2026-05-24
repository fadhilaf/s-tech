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
			c.Header("HX-Redirect", "/401")
			c.AbortWithStatusJSON(401, gin.H{"error": "Sesi tidak valid atau telah berakhir"})
			return
		}
	}
}

func ShouldAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin := utils.GetAdminFromContext(c)

		if !isAdmin {
			c.Header("HX-Redirect", "/401")
			c.AbortWithStatusJSON(401, gin.H{"error": "Kamu harus admin untuk mengakses halaman ini"})
			return
		}
	}
}

func ShouldAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := utils.GetUserIdFromContext(c)
		isAdmin := utils.GetAdminFromContext(c)

		if userId == uuid.Nil && !isAdmin {
			c.Header("HX-Redirect", "/login")
			c.AbortWithStatusJSON(401, gin.H{"error": "Sesi tidak valid atau telah berakhir"})
			return
		}
	}
}
