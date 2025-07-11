package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"task04/internal/common"
	"task04/internal/common/response"
)

func CheckJWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == " " {
			response.CustomFailMsg(http.StatusUnauthorized, "未登录", c)
		}
		if !strings.HasPrefix(authorization, "Bearer ") {
			response.CustomFailMsg(http.StatusUnauthorized, "未登录", c)
		}
		authArray := strings.SplitN(authorization, " ", 2)
		tokenString := authArray[1]
		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				response.CustomFailMsg(http.StatusUnauthorized, "未登录", c)
			}
			return []byte("123456"), nil
		})
		if err != nil {
			response.CustomFailMsg(http.StatusUnauthorized, "未登录", c)
		}
		if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
			username := (*claims)["username"].(string)
			id := int32((*claims)["id"].(float64))
			current := common.CurrentUserSession{Id: id, UserName: username}
			c.Set("CurrentUserSession", current)
			c.Next()
		} else {
			response.CustomFailMsg(http.StatusUnauthorized, "未登录", c)
		}
		return
	}
}
