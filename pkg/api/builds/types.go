// Package builds relates to Jenkins builds
// A Jenkins build is an execution of a particular job
package builds

import (
	"strconv"
	"time"

	durafmt "github.com/hako/durafmt"
)

// BuildResult is an enum of possible Jenkins build results
type BuildResult string

// Possible BuildResults
const (
	Success BuildResult = "SUCCESS"
	Failure BuildResult = "FAILURE"
)

// Build is the execution of a Jenkins job
type Build struct {
	Class       string        `json:"_class"`
	Actions     []BuildAction `json:"actions"`
	ID          int           `json:"number"`
	Result      BuildResult   `json:"result"`
	Description string        `json:"description"`
	Timestamp   int64         `json:"timestamp"`
	Duration    int64         `json:"duration"`
}

// BuildAction is an interface for hudson.model.*Action
type BuildAction interface{}

// ParametersAction represents a set of parameters used to call a job.ParametersAction
// Maps to hudson.model.ParametersAction
type ParametersAction struct {
	Class      string           `json:"_class"`
	Parameters []BuildParameter `json:"parameters"`
}

// BuildParameter represents a hudson.model.*ParameterValue
type BuildParameter struct {
	Class string `json:"_class"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Builds is an alias for a slice of Build
// Implements output.Printable
type Builds []Build

// Headers returns the headers in the default output
// Used to implement output.Printable
func (b *Builds) Headers() []string {
	return []string{
		"Build ID",
		"Result",
		"Time",
		"Duration",
	}
}

// Rows returns the rows for the default output format
// Used to implement output.Printable
func (b *Builds) Rows() [][]string {
	rows := [][]string{}
	for _, build := range *b {
		formattedTime, err := durafmt.ParseString((time.Duration(build.Duration) * time.Millisecond).String())
		if err != nil {
			panic(err) // TODO: Don't panic this
		}
		rows = append(rows, []string{
			strconv.Itoa(build.ID),
			string(build.Result),
			time.Unix(build.Timestamp/1000, 0).String(),
			formattedTime.String(),
		})
	}
	return rows
}
