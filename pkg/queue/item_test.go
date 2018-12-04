package queue

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ItemIsExecuting(t *testing.T) {
	dat, err := ioutil.ReadFile("./sample_item.json")
	if err != nil {
		panic(err)
	}
	var item Item
	json.Unmarshal(dat, &item)
	assert.True(t, item.Executing())
}
