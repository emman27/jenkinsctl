package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var a = Artifacts{
	Artifact{
		DisplayPath:  "hello.world",
		FileName:     "hello.world",
		RelativePath: "/something/hello.world",
	},
}

func TestArtifacts_Headers_Row_Lengths(t *testing.T) {
	assert.Equal(t, len(a.Rows()[0]), len(a.Headers()))
}

func TestArtifactsJSON(t *testing.T) {
	assert.Equal(t, string(a.JSON()), "[{\"display_path\":\"hello.world\",\"file_name\":\"hello.world\",\"relative_path\":\"/something/hello.world\"}]")
}
