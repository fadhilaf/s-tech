package utils

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveUserToSession(c *gin.Context, id uuid.UUID) {
	session := sessions.Default(c)

	idString := id.String()

	session.Set("user_id", idString)

	session.Options(sessions.Options{
		// dikasi "/" biar artinyo accessible di seluruh route didalem "/" (artiny segalo route), defaultny dio di /auth/login gr gr kito buatny pas lagi di route itu
		Path:     "/",
		HttpOnly: true,
		MaxAge:   24 * 3600,
	})

	err := session.Save()
	if err != nil {
		fmt.Println("error session save:", err)
	}
}

func GetUserIdFromSession(c *gin.Context) uuid.UUID {
	session := sessions.Default(c)

	userId := session.Get("user_id")
	if userId != nil {
		parsedUserId, err := uuid.Parse(userId.(string))

		if err == nil {
			return parsedUserId

		}
	}
	return uuid.Nil
}

func GetUserIdFromContext(c *gin.Context) uuid.UUID {
	userId, exist := c.Get("user_id")
	if exist {
		parsedUserId, ok := userId.(uuid.UUID)
		if ok {
			return parsedUserId
		}
	}
	return uuid.Nil
}

func SaveAdminToSession(c *gin.Context, isAdmin bool) {
	session := sessions.Default(c)
	session.Set("is_admin", isAdmin)
	session.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   24 * 3600,
	})
	err := session.Save()
	if err != nil {
		fmt.Println("error session save:", err)
	}
}

func GetAdminFromSession(c *gin.Context) bool {
	session := sessions.Default(c)

	isAdmin := session.Get("is_admin")

	if isAdmin == nil {
		return false
	} else {
		return isAdmin.(bool)
	}
}

func GetAdminFromContext(c *gin.Context) bool {
	isAdmin, exist := c.Get("is_admin")
	if exist {
		return isAdmin.(bool)
	}
	return false
}

func RemoveAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("user_id") != nil {
		session.Delete("user_id")
	}

	if session.Get("is_admin") != nil {
		session.Delete("is_admin")
	}

	session.Save()
}
