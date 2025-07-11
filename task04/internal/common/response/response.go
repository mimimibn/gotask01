package response

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: SUCCESS, Data: make(map[string]string), Msg: "操作成功"})
}

func FailMsg(msg string, c *gin.Context) {
	log.Println()
	c.JSON(http.StatusInternalServerError, Response{Code: ERROR, Data: make(map[string]string), Msg: msg})
}

func Data(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: SUCCESS, Data: data, Msg: ""})
}
func CustomFailMsg(code int, msg string, c *gin.Context) {
	c.JSON(code, Response{Code: SUCCESS, Data: make(map[string]string), Msg: msg})
}
