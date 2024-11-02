package linter1

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "linter1",
	Doc:      "linter1はすごい静的解析ツールです",
	Run:      run,
	Requires: []*analysis.Analyzer{},
}

func run(pass *analysis.Pass) (any, error) {
	fmt.Println("hello world!")
	return nil, nil
}
