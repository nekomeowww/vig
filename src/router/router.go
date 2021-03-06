package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/nekomeowww/vig/src/config"
	"github.com/nekomeowww/vig/src/controllers"
	"github.com/nekomeowww/vig/src/handler"
	"github.com/nekomeowww/vig/src/middleware"
)

// InitCORS 初始化跨域配置
func InitCORS(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Cookie", "Authorization", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		ExposeHeaders:    nil,
	}))
	return
}

// InitRouter 初始化主机模式路由
func InitRouter() *gin.Engine {
	gin.SetMode(config.Stage)
	r := gin.Default()

	// Static
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))
	r.Use(middleware.FrontendHandler())

	api := r.Group("/api")
	v1 := api.Group("v1")

	// ping

	v1.GET("ping", handler.NewHandler(controllers.ActionPing))

	// CORS
	InitCORS(r)

	return r
}
