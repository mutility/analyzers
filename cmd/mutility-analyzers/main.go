package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/mutility/unused-results/unret"
	"github.com/mutility/variable-format/varfmt"
)

func main() {
	multichecker.Main(
		varfmt.Analyzer().Analyzer,
		unret.Analyzer().Analyzer,
	)
}
