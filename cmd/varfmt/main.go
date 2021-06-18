package main

import (
	"github.com/mutility/analyzers/varfmt"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(varfmt.Analyzer)
}
