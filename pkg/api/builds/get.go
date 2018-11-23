package builds

import (
	"fmt"
	"io/ioutil"

	"github.com/emman27/jenkinsctl/pkg/api"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// List gets a list of builds for the specified job name
func List(c *api.JenkinsClient, jobName string) (*Builds, error) {
	resp, err := c.Get(fmt.Sprintf("%s/job/%s/api/json", c.BaseURL, jobName))
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get jobs for %s", jobName)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not read response body")
	}
	glog.Info(string(content))
	return &Builds{}, nil
}
