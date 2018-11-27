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

func Test_ParameterMarshalJSON(t *testing.T) {
	p := BuildParameter{
		Class: Boolean,
		Name:  "some Name",
		Value: "Hi",
	}
	json, err := p.MarshalJSON()
	assert.Nil(t, err, "Cannot marshal BuildParameter to JSON")
	assert.Equal(t, "{\"name\":\"some Name\",\"type\":\"Boolean\",\"value\":\"Hi\"}", string(json), "JSON result not as expected")
}
