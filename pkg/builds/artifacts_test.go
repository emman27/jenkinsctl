package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtifacts_Headers_Row_Lengths(t *testing.T) {
	a := Artifacts{
		Artifact{
			DisplayPath:  "hello.world",
			FileName:     "hello.world",
			RelativePath: "/something/hello.world",
		},
	}
	assert.Equal(t, len(a.Rows()[0]), len(a.Headers()))
}
