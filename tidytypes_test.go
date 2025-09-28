package tidytypes

import (
	"testing"

	"github.com/kisielk/errcheck/errcheck"
	"golang.org/x/tools/go/analysis/analysistest"
	"lesiw.io/checker"
	"lesiw.io/linelen"
)

func TestCheck(t *testing.T) {
	checker.Run(t,
		Analyzer,
		errcheck.Analyzer,
		linelen.Analyzer,
	)
}

func TestAnalysisTest(t *testing.T) {
	testdata := analysistest.TestData()
	pkgs := []string{"basic", "parameters", "returns"}
	for _, pkg := range pkgs {
		t.Run(pkg, func(t *testing.T) {
			analysistest.Run(t, testdata, Analyzer, pkg)
		})
	}
	t.Run("fix", func(t *testing.T) {
		analysistest.RunWithSuggestedFixes(t, testdata, Analyzer, "fix")
	})
}
