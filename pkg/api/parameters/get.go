package parameters

import (
	"errors"

	"github.com/emman27/jenkinsctl/pkg/api"
	"github.com/emman27/jenkinsctl/pkg/api/builds"
	"github.com/golang/glog"
)

const parameterActionClass = "hudson.model.ParametersAction"

// Get the parameters for a particular job and build
func Get(c *api.JenkinsClient, jobName string, buildID int) (*Parameters, error) {
	var (
		build *builds.Build
		err   error
	)
	if build, err = builds.Get(c, jobName, buildID); err != nil {
		return nil, err
	}
	glog.Infof("Build Actions found: %v", build.Actions)
	for _, action := range build.Actions {
		if action.Class == parameterActionClass {
			params := Parameters(*action.Parameters)
			return &params, nil
		}
	}
	return nil, errors.New("this build is not parameterized")
}
