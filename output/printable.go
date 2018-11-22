package output

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Printable is any object that can be printed to the command line
type Printable interface {
	// Headers are the headers to the default output format
	Headers() []string
	// Row is an array of strings to match to the headers
	Rows() [][]string
}

func Print(p Printable) {
	switch *outputFormat {
	default:
		printDefault(p)
	}
}

func printDefault(p Printable) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, strings.Join(p.Headers(), "\t"))
	for _, row := range p.Rows() {
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	w.Flush()
}
