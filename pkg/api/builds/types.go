// Package builds relates to Jenkins builds
// A Jenkins build is an execution of a particular job
package builds

import "fmt"

// Build is the execution of a Jenkins job
type Build struct {
	JobName     string
	BuildID     int
	Description string
}

// Builds is an alias for a slice of Build
// Implements output.Printable
type Builds []Build

// Headers returns the headers in the default output
// Used to implement output.Printable
func (b *Builds) Headers() []string {
	return []string{
		"Job Name",
		"Build ID",
		"Description",
	}
}

// Rows returns the rows for the default output format
// Used to implement output.Printable
func (b *Builds) Rows() [][]string {
	rows := [][]string{}
	for _, build := range *b {
		rows = append(rows, []string{build.JobName, fmt.Sprintf("%d", build.BuildID), build.Description})
	}
	return rows
}
