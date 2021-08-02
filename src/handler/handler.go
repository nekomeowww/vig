package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 状态码
const (
	Success = iota // 成功
)

// Response 响应
type Response interface{}

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// ErrResponse 错误响应
type ErrResponse struct {
	Status  int         `json:"-"`
	Code    int         `json:"code"`
	Message interface{} `json:"messsage"`
}

// Func 处理函数
type Func func(c *gin.Context) (Response, *ErrResponse)

// NewHandler 新建控制器
func NewHandler(handler Func) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, err := handler(c)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, response{Code: Success, Data: r})
	}
}
