package parameters

import (
	"errors"

	"github.com/emman27/jenkinsctl/pkg/api"
	"github.com/emman27/jenkinsctl/pkg/api/builds"
)

const parameterActionClass = "hudson.model.ParametersAction"

// Get the parameters for a particular job and build
func Get(c *api.JenkinsClient, jobName string, buildID int) (*Parameters, error) {
	build, err := builds.Get(c, jobName, buildID)
	if err != nil {
		return nil, err
	}
	return getParams(build)
}

func getParams(build *builds.Build) (*Parameters, error) {
	for _, action := range build.Actions {
		if action.Class == parameterActionClass {
			params := Parameters(*action.Parameters)
			return &params, nil
		}
	}
	return nil, errors.New("this build is not parameterized")
}
