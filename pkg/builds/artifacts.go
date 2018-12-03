package builds

import (
	"encoding/json"
)

// Artifact is a build artifact from Jenkins
type Artifact struct {
	DisplayPath  string `json:"display_path"`
	FileName     string `json:"file_name"`
	RelativePath string `json:"relative_path"`
}

// Artifacts is a collection of Artifacts.
// Implements output.Printable
type Artifacts []Artifact

// Headers for the default view
func (a *Artifacts) Headers() []string {
	return []string{
		"File Name",
		"Path",
	}
}

// Rows for the default view
func (a *Artifacts) Rows() [][]string {
	rows := [][]string{}
	for _, artifact := range *a {
		rows = append(rows, []string{artifact.FileName, artifact.RelativePath})
	}
	return rows
}

// JSON converts to JSON
func (a *Artifacts) JSON() []byte {
	res, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return res
}
