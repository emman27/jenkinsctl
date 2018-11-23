package api

import (
	"net/http"
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
