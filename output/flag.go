package output

import (
	"flag"
)

var outputFormat = flag.String("o", "", "Output format to use. Choose one of default (empty string), wide, json or yaml")
