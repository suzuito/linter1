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
	// 概要説明
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
	// passという変数には、構文解析済みのASTの情報、型の情報が格納されています。
	// 開発者は、passに格納されている情報を用いることで、静的解析処理を記述します。

	// inspect.Analyzerの実行結果は、ASTを走査する機能を提供する。
	// ASTを走査する処理を自前で書くのはダルいので、そのまま利用させてもらう。
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// ASTの走査
	inspect.Preorder(
		// 第1引数に指定したノードを訪問したときに
		// 第2引数の関数が実行される
		[]ast.Node{
			(*ast.ValueSpec)(nil),
			(*ast.Field)(nil),
			(*ast.AssignStmt)(nil),
		},
		func(current ast.Node) {
			switch n := current.(type) {
			case *ast.ValueSpec:
				// const, var による宣言文で利用されている識別子を取得し、check関数へ
				for i, value := range n.Values {
					utyp := pass.TypesInfo.TypeOf(value).Underlying()
					check(pass, n.Names[i], utyp)
				}
			case *ast.Field:
				// struct, interfaceのメンバー、もしくは、
				// 関数の引数や返り値パラメータで利用されている識別子を取得し、check関数へ
				utyp := pass.TypesInfo.TypeOf(n.Type).Underlying()
				check(pass, n.Names[0], utyp)
			case *ast.AssignStmt:
				// := による変数割当ての左側にある識別子を取得し、check関数へ
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
		},
	)

	return nil, nil
}

// 識別子identの型typがブール型だった場合に、識別子名がpatternに一致するかどうかをチェックする
func check(
	pass *analysis.Pass,
	ident *ast.Ident,
	typ types.Type,
) {
	if types.Identical(typ, types.Typ[types.Bool]) ||
		types.Identical(typ, types.Typ[types.UntypedBool]) {
		if matched := pattern.MatchString(ident.Name); !matched {
			// pass変数のReportf関数を用いることで、警告文をいい感じのフォーマットで出力できます。
			pass.Reportf(
				ident.Pos(),
				"a boolean variable does not match pattern",
			)
		}
	}
}
