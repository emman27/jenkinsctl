package builds

// Artifact is a build artifact from Jenkins
type Artifact struct {
	DisplayPath  string
	FileName     string
	RelativePath string
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
// TODO: stub
func (a *Artifacts) JSON() []byte {
	return []byte{}
}
