package linter2

import (
	"go/ast"
	"go/types"
	"regexp"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// 静的解析処理の定義
var Analyzer = &analysis.Analyzer{
	// 名前
	Name: "linter2",
	// 概要説明(どこで使われてるのかよくわからん)
	Doc: `linter2 is boolean variable naming checker`,
	// 静的解析処理をする関数(開発者の主な仕事はこの関数を作ること！)
	Run: run,
	// この静的解析処理が依存する別の静的解析処理
	// go/analyticsの目的の1つは、静的解析処理の再利用性を高めること
	// ある静的解析処理の結果を別の静的解析処理でも利用できる仕組みがRequires
	// ここに指定した静的解析処理の結果を、analysis.Pass.ResultOf変数から取り出して利用できる
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var pattenString = "^(is|has|are)"
var pattern = regexp.MustCompile(pattenString)

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ValueSpec)(nil),
		(*ast.Field)(nil),
		(*ast.AssignStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ValueSpec:
			for i, value := range n.Values {
				utyp := pass.TypesInfo.TypeOf(value).Underlying()
				check(pass, n.Names[i], utyp)
			}
		case *ast.Field:
			utyp := pass.TypesInfo.TypeOf(n.Type).Underlying()
			check(pass, n.Names[0], utyp)
		case *ast.AssignStmt:
			for i := range n.Lhs {
				l, r := n.Lhs[i], n.Rhs[i]
				lIdent, ok := l.(*ast.Ident)
				if !ok {
					panic(ok)
				}
				utyp := pass.TypesInfo.TypeOf(r).Underlying()
				check(pass, lIdent, utyp)
			}
		}
	})

	return nil, nil
}

func check(
	pass *analysis.Pass,
	name *ast.Ident,
	valueType types.Type,
) {
	if types.Identical(valueType, types.Typ[types.Bool]) ||
		types.Identical(valueType, types.Typ[types.UntypedBool]) {
		if matched := pattern.MatchString(name.Name); !matched {
			pass.Reportf(
				name.Pos(),
				"a boolean variable does not match pattern",
			)
		}
	}
}
