package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emman27/jenkinsctl/pkg/queue"
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

func Test_GetQueueItemExecution(t *testing.T) {
	i := 0
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		dat, err := ioutil.ReadFile("./sample_item.json")
		assert.Nil(t, err)
		var item queue.Item
		err = json.Unmarshal(dat, &item)
		assert.Nil(t, err)
		if i < 2 {
			item.Executable = nil
		}
		body, err := json.Marshal(item)
		assert.Nil(t, err)
		res.Write(body)
		i++
	})
	server := httptest.NewServer(handler)
	c := NewJenkinsClient(server.URL, "", "")
	execution, err := c.GetQueueItemExecution(71)
	assert.Nil(t, err)
	assert.NotNil(t, execution)
}

func Test_GetQueueItem404(t *testing.T) {
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method)
		res.WriteHeader(http.StatusNotFound)
	})
	server := httptest.NewServer(handler)
	c := NewJenkinsClient(server.URL, "", "")
	item, err := c.GetQueueItem(71)
	assert.NotNil(t, err)
	assert.Nil(t, item)
}
