package parameters

import "github.com/emman27/jenkinsctl/pkg/api/builds"

// Parameters is the printable array of Parameter
// Implements output.Printable
type Parameters []builds.BuildParameter

// Headers for the default view
func (p *Parameters) Headers() []string {
	return []string{
		"Type",
		"Name",
		"Value",
	}
}

// Rows for the default view
func (p *Parameters) Rows() [][]string {
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
func (p *Parameters) JSON() []byte {
	return []byte("")
}
