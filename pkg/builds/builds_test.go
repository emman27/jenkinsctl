package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilds_Headers_Rows_Length(t *testing.T) {
	builds := Builds{
		{ID: 2, Result: Success},
	}
	headers := new(Builds).Headers()
	assert.Equal(t, len(headers), len(builds.Rows()[0]), "Headers and Rows have different lengths")
}

func TestBuilds_JSON(t *testing.T) {
	builds := Builds{
		{
			Class:       "Something",
			Actions:     []BuildAction{},
			ID:          24,
			Result:      Success,
			Description: "Just a job",
			Timestamp:   int64(123456789),
			Duration:    int64(30),
		},
	}
	assert.Equal(t, "[{\"_class\":\"Something\",\"actions\":[],\"number\":24,\"result\":\"SUCCESS\",\"description\":\"Just a job\",\"timestamp\":123456789,\"duration\":30,\"artifacts\":null}]", string(builds.JSON()))
}

func Test_GenerateParameters(t *testing.T) {
	result, err := GenerateParametersBody(map[string]interface{}{
		"hello": "world",
	})
	assert.Nil(t, err)
	assert.Equal(t, "json=%7B%22parameter%22%3A%5B%7B%22name%22%3A%22hello%22%2C%22value%22%3A%22world%22%7D%5D%7D", result)
}
