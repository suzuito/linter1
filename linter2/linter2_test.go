package linter2_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/suzuito/linter1/linter2"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, linter2.Analyzer, "a")
}
