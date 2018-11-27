package parameters

import (
	"testing"

	"github.com/emman27/jenkinsctl/pkg/builds"
	"github.com/stretchr/testify/assert"
)

func TestHeaders_and_Rows_Length_Equal(t *testing.T) {
	p := Parameters{
		builds.BuildParameter{
			Class: builds.Boolean,
			Name:  "some Name",
			Value: "Hi",
		},
	}
	assert.Equal(t, len(p.Headers()), len(p.Rows()[0]))
}
