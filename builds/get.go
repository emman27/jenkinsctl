package builds

import "fmt"

type Build struct {
	JobName     string
	BuildID     int
	Description string
}

type Builds []Build

func (b *Builds) Headers() []string {
	return []string{
		"Job Name",
		"Build ID",
		"Description",
	}
}

func (b *Builds) Rows() [][]string {
	rows := [][]string{}
	for _, build := range *b {
		rows = append(rows, []string{build.JobName, fmt.Sprintf("%d", build.BuildID), build.Description})
	}
	return rows
}

func List() *Builds {
	return &Builds{
		{JobName: "Hello", BuildID: 1, Description: "test"},
	}
}
