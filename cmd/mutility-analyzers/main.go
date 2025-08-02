package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/mutility/parameter-swap/pswap"
	"github.com/mutility/undeferred/undefer"
	"github.com/mutility/unused-results/unret"
	"github.com/mutility/variable-format/varfmt"
)

func main() {
	multichecker.Main(
		varfmt.Analyzer().Analyzer,
		unret.Analyzer().Analyzer,
		pswap.Analyzer().Analyzer,
		undefer.Analyzer().Analyzer,
	)
}
