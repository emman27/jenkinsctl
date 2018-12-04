package api

import (
	"io/ioutil"
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

func Test_CreateBuildWithAPIFailure(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/job/my-job/build", req.URL.Path)
		assert.Equal(t, "POST", req.Method)
		res.WriteHeader(http.StatusBadRequest)
	})
	server := httptest.NewServer(handler)
	c := NewJenkinsClient(server.URL, "user", "password")
	_, err := c.CreateBuild("my-job", map[string]interface{}{})
	assert.NotNil(t, err)
}

func Test_CreateBuildWithParams(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/job/my-job/build", req.URL.Path)
		assert.Equal(t, "POST", req.Method)
		dat, err := ioutil.ReadAll(req.Body)
		assert.Nil(t, err)
		assert.Equal(t, "json=%7B%22parameter%22%3A%5B%7B%22name%22%3A%22hello%22%2C%22value%22%3A%22world%22%7D%5D%7D", string(dat))
	})
	server := httptest.NewServer(handler)
	c := NewJenkinsClient(server.URL, "user", "password")
	_, err := c.CreateBuild("my-job", map[string]interface{}{
		"hello": "world",
	})
	assert.Nil(t, err)
}
