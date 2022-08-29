package main

import (
	"github.com/mutility/variable-format/varfmt"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(varfmt.Analyzer().Analyzer)
}
