package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/emman27/jenkinsctl/pkg/queue"
	"github.com/pkg/errors"
)

// GetQueueItem retrieves a queue item
func (c *JenkinsClient) GetQueueItem(ID int) (*queue.Item, error) {
	resp, err := c.Get(fmt.Sprintf("/queue/item/%d/api/json", ID))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch item from queue")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read response body")
	}
	var item queue.Item
	err = json.Unmarshal(body, &item)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot convert JSON to Item")
	}
	return &item, nil
}

// GetQueueItemExecution will wait until a queue item is executed, then return that Execution
func (c *JenkinsClient) GetQueueItemExecution(ID int) (*queue.Executable, error) {
	var item = &queue.Item{}
	var err error
	for !item.Executing() {
		item, err = c.GetQueueItem(ID)
		if err != nil {
			return nil, err
		}
		time.Sleep(500 * time.Millisecond)
	}
	return item.Executable, nil
}
