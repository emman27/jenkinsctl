package builds

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/emman27/jenkinsutils/pkg/api"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// List gets a list of builds for the specified job name
func List(c *api.JenkinsClient, jobName string) (*Builds, error) {
	resp, err := c.Get(fmt.Sprintf("%s/job/%s/api/json", c.BaseURL, jobName))
	if err != nil {
		return nil, errors.Wrapf(err, "Could not fetch builds for job %s", jobName)
	}
	if resp.StatusCode != http.StatusOK {
		glog.Warning(resp.Request.Header)
		return nil, fmt.Errorf("failed to fetch builds for job %s, got HTTP status code %d", jobName, resp.StatusCode)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not read response body")
	}
	glog.Info(string(content))
	return &Builds{}, nil
}
