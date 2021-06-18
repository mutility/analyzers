package varfmt_test

import (
	"testing"

	"github.com/mutility/analyzers/varfmt"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, varfmt.Analyzer, "a")
}
