package builds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/emman27/jenkinsctl/pkg/api"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// List gets a list of builds for the specified job name
func List(c *api.JenkinsClient, jobName string) (*Builds, error) {
	resp, err := c.Get(fmt.Sprintf("/job/%s/api/json", jobName))
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

// Get retrieves a particular build of a job
func Get(c *api.JenkinsClient, jobName string, buildID int) (*Build, error) {
	resp, err := c.Get(fmt.Sprintf("/job/%s/%d/api/json", jobName, buildID))
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get job %d for %s", buildID, jobName)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not read response body")
	}
	var build = new(Build)
	json.Unmarshal(content, build)
	glog.Infof("Parsed response %v", build)
	return build, nil
}
