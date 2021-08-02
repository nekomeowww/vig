package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewHandler(t *testing.T) {
	h := func(c *gin.Context) (resp Response, err *ErrResponse) {
		return nil, nil
	}
	NewHandler(h)

	h = func(c *gin.Context) (Response, *ErrResponse) {
		return nil, &ErrResponse{}
	}
	NewHandler(h)
}
