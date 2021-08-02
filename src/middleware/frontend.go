package middleware

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"

	// Static Files
	"github.com/nekomeowww/vig/src/logger"
	_ "github.com/nekomeowww/vig/statik"
)

// FrontendHandler 前端静态文件处理
func FrontendHandler() gin.HandlerFunc {
	statikFS, err := fs.New()
	if err != nil {
		logger.Fatal(err)
	}

	// 读取index.html
	r, err := statikFS.Open("/index.html")
	if err != nil {
		logger.Fatal(err)
	}

	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		logger.Fatal(err)
	}

	fileContent := string(contents)
	fileServer := http.FileServer(statikFS)
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// API 跳过
		if strings.HasPrefix(path, "/api") || path == "/manifest.json" {
			c.Next()
			return
		}

		// 不存在的路径和index.html均返回index.html
		_, err := statikFS.Open("/" + path)
		if (path == "/index.html") || (path == "/") || err != nil {
			c.Header("Content-Type", "text/html")
			c.String(200, fileContent)
			c.Abort()
			return
		}

		// 存在的静态文件
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
