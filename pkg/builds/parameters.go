package builds

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

// Parameters is the printable array of Parameter
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
	return []byte("")
}
