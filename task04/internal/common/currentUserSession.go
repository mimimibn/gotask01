package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task04/internal/common/response"
)

type CurrentUserSession struct {
	Id       int32
	UserName string
}

func GetCurrentUserSession(c *gin.Context) (currentUserSession CurrentUserSession) {
	value, exists := c.Get("CurrentUserSession")
	if !exists {
		response.CustomFailMsg(http.StatusUnauthorized, "未登录", c)
		return
	}
	session := value.(CurrentUserSession)
	return session
}
