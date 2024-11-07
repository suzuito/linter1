package main

import (
	"github.com/suzuito/linter1/linter2"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(linter2.Analyzer) }
