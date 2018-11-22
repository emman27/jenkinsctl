package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewJenkinsClient("google.com", "hello", "world")
	assert.Equal(t, "Basic aGVsbG86d29ybGQ=", client.authorizationHeader, "Base64 encoding is incorrect")
}
