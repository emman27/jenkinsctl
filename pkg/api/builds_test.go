package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
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
	_, err := c.CreateBuild("my-job", map[string]string{})
	assert.Nil(t, err)
}

func Test_CreateBuildWithAPIFailure(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/job/my-job/build", req.URL.Path)
		assert.Equal(t, "POST", req.Method)
		res.WriteHeader(http.StatusBadRequest)
	})
	server := httptest.NewServer(handler)
	c := NewJenkinsClient(server.URL, "user", "password")
	_, err := c.CreateBuild("my-job", map[string]string{})
	assert.NotNil(t, err)
}

func Test_CreateBuildWithParams(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/job/my-job/buildWithParameters", req.URL.Path)
		assert.Equal(t, url.Values{"hello": []string{"world"}}, req.URL.Query())
		assert.Equal(t, "POST", req.Method)
	})
	server := httptest.NewServer(handler)
	c := NewJenkinsClient(server.URL, "user", "password")
	_, err := c.CreateBuild("my-job", map[string]string{
		"hello": "world",
	})
	assert.Nil(t, err)
}
