package parameters

import (
	"testing"

	"github.com/emman27/jenkinsctl/pkg/builds"
	"github.com/stretchr/testify/assert"
)

func Test_getParamsErrorIfNoParams(t *testing.T) {
	build := builds.Build{
		Actions: []builds.BuildAction{},
	}
	_, err := getParams(&build)
	assert.NotNil(t, err)
}

func Test_getParams(t *testing.T) {
	build := builds.Build{
		Actions: []builds.BuildAction{
			{Class: parameterActionClass, Parameters: &[]builds.BuildParameter{}},
		},
	}
	params, err := getParams(&build)
	assert.Nil(t, err)
	assert.NotNil(t, params)
}
