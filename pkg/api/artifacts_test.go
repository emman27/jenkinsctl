package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetArtifacts(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		dat, err := ioutil.ReadFile("./sample_build.json")
		assert.Nil(t, err)
		res.Write(dat)
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	client := NewJenkinsClient(server.URL, "user", "password")
	artifacts, err := client.GetArtifacts("my-job", 24)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(*artifacts))
}

func Test_GetArtifactsNotFound(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNotFound)
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	client := NewJenkinsClient(server.URL, "user", "password")
	_, err := client.GetArtifacts("my-job", 24)
	assert.NotNil(t, err)
}
