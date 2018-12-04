package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateBuildWithoutParams(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/job/my-job/build", req.URL.Path)
		assert.Equal(t, "POST", req.Method)
	})
	server := httptest.NewServer(handler)
	c := NewJenkinsClient(server.URL, "user", "password")
	_, err := c.CreateBuild("my-job", map[string]interface{}{})
	assert.Nil(t, err)
}
