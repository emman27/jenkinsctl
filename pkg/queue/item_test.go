package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ItemIsExecuting(t *testing.T) {
	item := Item{
		Executable: &Executable{
			Number: 5,
			URL:    "https://jenkins.com/executing",
		},
	}
	assert.True(t, item.Executing())
}
