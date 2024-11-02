package main

import (
	"github.com/suzuito/linter1/linter1"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(linter1.Analyzer)
}
