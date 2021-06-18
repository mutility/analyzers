package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/mutility/analyzers/varfmt"
)

func main() {
	multichecker.Main(
		varfmt.Analyzer,
	)
}
