package api

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewJenkinsClient("google.com", "hello", "world")
	assert.Equal(t, "Basic aGVsbG86d29ybGQ=", client.authorizationHeader, "Base64 encoding is incorrect")
}

func Test_checkStatusCode(t *testing.T) {
	a := assert.New(t)
	a.NotNil(checkStatusCode(&http.Response{StatusCode: 400}), "HTTP Code 400 should cause errors")
	a.NotNil(checkStatusCode(&http.Response{StatusCode: 403}), "HTTP Code 403 should cause errors")
	a.NotNil(checkStatusCode(&http.Response{StatusCode: 500}), "HTTP Code 500 should cause errors")
	a.Nil(checkStatusCode(&http.Response{StatusCode: 200}), "HTTP Code 200 should not cause errors")
	a.Nil(checkStatusCode(&http.Response{StatusCode: 201}), "HTTP Code 201 should not cause errors")
}

func Test_DoAddsHeaders(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/", req.URL.EscapedPath())
		assert.Equal(t, "Basic ", req.Header.Get("Authorization")[:6])
		bytes, err := base64.StdEncoding.DecodeString(req.Header.Get("Authorization")[6:])
		assert.Nil(t, err)
		assert.Equal(t, "user:password", string(bytes))
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	client := NewJenkinsClient(server.URL, "user", "password")
	client.Get("/")
}

func Test_PostSendsBody(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		assert.Nil(t, err)
		assert.Equal(t, string(body), "test message")
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	client := NewJenkinsClient(server.URL, "user", "password")
	client.Post("/", strings.NewReader("test message"))
}
