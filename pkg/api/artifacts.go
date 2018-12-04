package api

import (
	"fmt"
	"io/ioutil"

	"github.com/emman27/jenkinsctl/pkg/builds"
)

// GetArtifacts returns an array of artifacts for a particular build
func (c *JenkinsClient) GetArtifacts(jobName string, buildID int) (*builds.Artifacts, error) {
	build, err := c.GetBuild(jobName, buildID)
	if err != nil {
		return nil, err
	}
	result := builds.Artifacts(build.Artifacts)
	return &result, nil
}

// GetArtifact retrieves an artifact from Jenkins.
// Returns a byte array containing the contents of the artifact
// Artifacts are typically files stored by the Jenkins job
func (c *JenkinsClient) GetArtifact(jobName string, buildID int, artifactFileName string) ([]byte, error) {
	artifacts, err := c.GetArtifacts(jobName, buildID)
	if err != nil {
		return []byte{}, err
	}
	for _, artifact := range *artifacts {
		if artifact.FileName == artifactFileName {
			return c.getArtifactContent(jobName, buildID, artifact.RelativePath)
		}
	}
	return []byte{}, fmt.Errorf("artifact %s does not exist for %s #%d", artifactFileName, jobName, buildID)
}

func (c *JenkinsClient) getArtifactContent(jobName string, buildID int, artifactRelativePath string) ([]byte, error) {
	resp, err := c.Get(fmt.Sprintf("/job/%s/%d/artifact/%s", jobName, buildID, artifactRelativePath))
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
