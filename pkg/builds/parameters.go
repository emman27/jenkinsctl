package builds

import "encoding/json"

// Possible Parameter Types
const (
	Boolean string = "hudson.model.BooleanParameterValue"
	String  string = "hudson.model.StringParameterValue"
	File    string = "hudson.model.FileParameterValue"
)

// Type gets a string representation of the type of the BuildParameter
// Supported types are booleans, strings and files
func (p *BuildParameter) Type() string {
	switch p.Class {
	case Boolean:
		return "Boolean"
	case String:
		return "String"
	case File:
		return "File"
	default:
		panic("Unknown parameter type. Please let the developers know!")
	}
}

// BuildParameter represents a hudson.model.*ParameterValue
type BuildParameter struct {
	Class string `json:"_class"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// BuildParameters is the printable array of Parameter
// Implements output.Printable
type BuildParameters []BuildParameter

// Headers for the default view
func (p *BuildParameters) Headers() []string {
	return []string{
		"Type",
		"Name",
		"Value",
	}
}

// Rows for the default view
func (p *BuildParameters) Rows() [][]string {
	result := [][]string{}
	for _, param := range *p {
		result = append(result, []string{
			param.Type(),
			param.Name,
			param.Value,
		})
	}
	return result
}

// JSON formatted parameters
// TODO: STUB
func (p *BuildParameters) JSON() []byte {
	res, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return res
}

// MarshalJSON provides a custom JSON output format for a BuildParameter
// This means you get {"type": "Boolean"} instead of {"type": <some-java-class-name>}
// Makes things much nicer to work with
func (p *BuildParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":  p.Type(),
		"name":  p.Name,
		"value": p.Value,
	})
}
