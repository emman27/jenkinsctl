package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaders_and_Rows_Length_Equal(t *testing.T) {
	p := BuildParameters{
		BuildParameter{
			Class: Boolean,
			Name:  "some Name",
			Value: "Hi",
		},
	}
	assert.Equal(t, len(p.Headers()), len(p.Rows()[0]))
}
