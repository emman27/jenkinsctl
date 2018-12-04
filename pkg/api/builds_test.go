package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getBuildsTestServer() *httptest.Server {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			res.Header().Set("Location", "https://jenkins.com/job/my-job/454/")
		} else if req.Method == "GET" && strings.HasPrefix(req.URL.Path, "/job") {
			dat, err := ioutil.ReadFile("./sample_build.json")
			if err != nil {
				panic(err)
			}
			res.Write(dat)
		} else if req.Method == "GET" && strings.HasPrefix(req.URL.Path, "/queue") {
			dat, err := ioutil.ReadFile("./sample_item.json")
			if err != nil {
				panic(err)
			}
			res.Write(dat)
		} else {
			panic("Not supported")
		}
	})
	return httptest.NewServer(handler)
}

func Test_CreateBuildWithoutParams(t *testing.T) {
	server := getBuildsTestServer()
	defer server.Close()
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
	defer server.Close()
	c := NewJenkinsClient(server.URL, "user", "password")
	_, err := c.CreateBuild("my-job", map[string]string{})
	assert.NotNil(t, err)
}

func Test_CreateBuildWithParams(t *testing.T) {
	server := getBuildsTestServer()
	defer server.Close()
	c := NewJenkinsClient(server.URL, "user", "password")
	_, err := c.CreateBuild("my-job", map[string]string{
		"hello": "world",
	})
	assert.Nil(t, err)
}
