package controllers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nekomeowww/vig/handler"
)

func TestActionPing(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	rec, c := handler.CreateTestContext(http.MethodGet, "", nil)
	r, errResp := ActionPing(c)
	require.Nil(errResp)
	require.NotEmpty(r)
	assert.Equal(200, rec.Code)

	resp := r.(*pingResp)
	assert.Equal("pong", resp.Message)
	assert.Equal("test", resp.Stage)
}
