package linter2_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/suzuito/linter1/linter2"
	"golang.org/x/tools/go/analysis/analysistest"
)

// 静的解析ツールのテストコード
// testdataディレクトリ配下にあるソースコードを入力として静的解析ツールを動かします

func TestAnalyzer(t *testing.T) {
	// この行がなにやってるかはあんまり詳しく見てない
	dirPathTestdata := testutil.WithModules(t, analysistest.TestData(), nil)
	// この行でテストを実行します
	analysistest.Run(t, dirPathTestdata, linter2.Analyzer, "a")
	analysistest.Run(t, dirPathTestdata, linter2.Analyzer, "b")

	// 期待値はテスト対象となっているGoのソースコードの中に記述されています。
	// testdata/src/a/a.go を参照してください。
}
