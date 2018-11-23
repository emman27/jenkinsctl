package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilds_Headers_Rows_Length(t *testing.T) {
	builds := Builds{
		{JobName: "Test Job", BuildID: 2, Description: "Hello World"},
	}
	headers := new(Builds).Headers()
	assert.Equal(t, len(headers), len(builds.Rows()[0]), "Headers and Rows have different lengths")
}
