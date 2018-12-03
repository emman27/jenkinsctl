package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetArtifacts(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/job/my-job/24/api/json", req.URL.Path)
		dat, err := ioutil.ReadFile("./sample_build.json")
		if err != nil {
			panic(err)
		}
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
		assert.Equal(t, "/job/my-job/24/api/json", req.URL.Path)
		res.WriteHeader(http.StatusNotFound)
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	client := NewJenkinsClient(server.URL, "user", "password")
	_, err := client.GetArtifacts("my-job", 24)
	assert.NotNil(t, err)
}

func Test_GetArtifactContent(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if strings.Contains(req.URL.Path, "artifact") {
			assert.Equal(t, "/job/my-job/24/artifact/something/hello.txt", req.URL.Path)
			res.Write([]byte("hello world"))
		} else {
			dat, err := ioutil.ReadFile("./sample_build.json")
			if err != nil {
				panic(err)
			}
			res.Write(dat)
		}
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	client := NewJenkinsClient(server.URL, "user", "password")
	artifact, err := client.GetArtifact("my-job", 24, "hello.txt")
	assert.Nil(t, err)
	assert.Equal(t, "hello world", string(artifact))
}
