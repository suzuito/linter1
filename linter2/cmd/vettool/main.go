// このソースコードをコンパイルしてできるバイナリ(例 vettool.outとする)は
// `go vet`コマンドから呼び出すことができる。(例 `go vet -vettool=vettool.out ./...`)。
// main関数の中で、引数にAnalyzerを指定してunitchecket.Main関数を呼ぶだけ。
package main

import (
	"github.com/suzuito/linter1/linter2"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(linter2.Analyzer)
}
