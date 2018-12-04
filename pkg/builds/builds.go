// Package builds relates to Jenkins builds
// A Jenkins build is an execution of a particular job
package builds

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	durafmt "github.com/hako/durafmt"
	"github.com/pkg/errors"
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
	Artifacts   []Artifact    `json:"artifacts"`
}

// BuildAction is an interface for hudson.model.*Action
type BuildAction struct {
	Class      string            `json:"_class"`
	Parameters *[]BuildParameter `json:"parameters"`
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

// JSON formats a build as JSON
func (b *Builds) JSON() []byte {
	res, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return res
}

// GenerateParametersBody converts a map of parameters into a Jenkins readable format
// See https://wiki.jenkins.io/display/JENKINS/Remote+access+API
func GenerateParametersBody(content map[string]interface{}) (string, error) {
	params := map[string][]map[string]interface{}{
		"parameter": []map[string]interface{}{},
	}
	for k, v := range content {
		params["parameter"] = append(params["parameter"], map[string]interface{}{
			"name":  k,
			"value": v,
		})
	}
	dat, err := json.Marshal(params)
	if err != nil {
		return "", errors.Wrap(err, "Could not convert content to JSON")
	}
	v := url.Values{}
	v.Add("json", string(dat))
	return v.Encode(), nil
}
