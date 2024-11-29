package main

import (
	"github.com/suzuito/linter1/linter2"
	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		linter2.Analyzer,
	}, nil
}
