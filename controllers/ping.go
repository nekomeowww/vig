package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/nekomeowww/vig/config"
	"github.com/nekomeowww/vig/handler"
)

type pingResp struct {
	Stage   string `json:"stage"`
	Message string `json:"message"`
}

// ActionPing ping
func ActionPing(c *gin.Context) (handler.Response, *handler.ErrResponse) {
	return &pingResp{
		Stage:   config.Stage,
		Message: "pong",
	}, nil
}
