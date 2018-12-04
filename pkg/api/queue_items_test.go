package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetQueueItem(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		dat, err := ioutil.ReadFile("./sample_item.json")
		assert.Nil(t, err)
		res.Write(dat)
	})
	server := httptest.NewServer(handler)
	c := NewJenkinsClient(server.URL, "", "")
	item, err := c.GetQueueItem(71)
	assert.Nil(t, err)
	assert.True(t, item.Executing())
}
