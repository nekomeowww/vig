package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/nekomeowww/vig/config"
)

// CreateTestContext 创建新的测试 context
func CreateTestContext(method, uri string, body interface{}) (*httptest.ResponseRecorder, *gin.Context) {
	recorder := httptest.NewRecorder()
	gin.SetMode(config.Stage)
	c, _ := gin.CreateTestContext(recorder)
	switch t := body.(type) {
	case nil:
		c.Request, _ = http.NewRequest(method, uri, nil)
		return recorder, c
	case io.Reader:
		c.Request, _ = http.NewRequest(method, uri, t)
		c.Request.Header.Set("Content-Type", "application/json")
		return recorder, c
	case string:
		query, err := url.ParseQuery(t)
		if err != nil {
			panic(err)
		}
		c.Request, _ = http.NewRequest(method, uri, nil)
		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		c.Request.PostForm = query
		return recorder, c
	case url.Values:
		c.Request, _ = http.NewRequest(method, uri, nil)
		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		c.Request.PostForm = t
		return recorder, c
	default:
		b, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		c.Request, _ = http.NewRequest(method, uri, bytes.NewBuffer(b))
		c.Request.Header.Set("Content-Type", "application/json")
		return recorder, c
	}
}
