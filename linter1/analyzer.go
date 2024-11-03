package linter1

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "linter1",
	Doc:  "linter1はすごい静的解析ツールです",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		&ast.Ident{},
	}
	for n := range inspect.PreorderSeq(nodeFilter...) {
		switch n := n.(type) {
		case *ast.Ident:
			t := pass.TypesInfo.TypeOf(n)
			fmt.Println("---", n, "---", t, "---")
		case *ast.AssignStmt:
		case *ast.DeclStmt:
		case *ast.FuncType:
		}

		// fmt.Println("---", nodeIdent, "---", nodeIdentType, "---", types.Identical(nodeIdentType.Type(), types.Typ[types.Bool]))

	}
	return nil, nil
}
