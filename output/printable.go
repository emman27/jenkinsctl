package output

import (
	"bytes"
	"encoding/json"
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
	JSON() []byte
}

// Print sends a message to stdout with the format set in the command line
func Print(p Printable) {
	switch *outputFormat {
	case "json":
		printJSON(p)
	default:
		printDefault(p)
	}
}

func printDefault(p Printable) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)
	fmt.Fprintln(w, strings.Join(p.Headers(), "\t"))
	for _, row := range p.Rows() {
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	w.Flush()
}

func printJSON(p Printable) {
	var out bytes.Buffer
	json.Indent(&out, p.JSON(), "", "  ")
	out.WriteTo(os.Stdout)
}
