package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/emman27/jenkinsctl/pkg/builds"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// GetBuild retrieves a particular build of a job
func (c *JenkinsClient) GetBuild(jobName string, buildID int) (*builds.Build, error) {
	resp, err := c.Get(fmt.Sprintf("/job/%s/%d/api/json", jobName, buildID))
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get job %d for %s", buildID, jobName)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not read response body")
	}
	var build = new(builds.Build)
	json.Unmarshal(content, build)
	glog.Infof("Parsed response %v", build)
	return build, nil
}

// CreateBuild starts a build in Jenkins
func (c *JenkinsClient) CreateBuild(jobName string, params map[string]string) (*builds.Build, error) {
	glog.Infof("Creating build %s with parameters %v", jobName, params)
	parameters, err := builds.GenerateParametersBody(params)
	glog.Infof("Parameters: %s", parameters)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to format parameters")
	}
	reader := strings.NewReader("")
	endpoint := fmt.Sprintf("/job/%s/build", jobName)
	// FIXME: Technically, this is wrong and it should be based on whether the job is parameterized or not.
	if parameters != "" {
		endpoint = fmt.Sprintf("/job/%s/buildWithParameters?%s", jobName, parameters)
	}
	resp, err := c.Post(endpoint, reader)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not create build")
	}
	glog.Infof("Queued build: %s", resp.Header.Get("Location"))
	return c.followLocationToBuild(resp.Header.Get("Location"))
}

func (c *JenkinsClient) followLocationToBuild(location string) (*builds.Build, error) {
	split := strings.Split(location, "/")
	queueItemID := split[len(split)-2]
	jobName := split[len(split)-3]
	queueItemIDInt, err := strconv.Atoi(queueItemID)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not parse a queue number")
	}
	execution, err := c.GetQueueItemExecution(queueItemIDInt)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get a matching execution for the job")
	}
	return c.GetBuild(jobName, execution.Number)

}
