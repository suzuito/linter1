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

	// 概要説明(正直、どこで使われてるのかよくわからん)
	Doc: `linter2 is boolean variable naming checker\n\n
like https://detekt.dev/docs/rules/naming/#booleanpropertynaming
`,

	// 静的解析処理をする関数(開発者の主な仕事はこの関数を作ること！)
	Run: run,

	// この静的解析処理が依存する別の静的解析処理
	// go/analyticsの目的の1つは、静的解析処理の再利用性を高めること。
	// Requiresは、ある静的解析処理の結果を別の静的解析処理でも利用できる仕組み。
	// ここに指定した静的解析処理の結果を、analysis.Pass.ResultOf変数から取り出して利用できる。
	Requires: []*analysis.Analyzer{
		// inspect.AnalyzerはASTを探索する機能(Inspector)を提供する
		// https://pkg.go.dev/golang.org/x/tools@v0.26.0/go/ast/inspector#Inspector
		// inspect.Analyzerのように
		// 他のAnalyzerから利用されることを前提とするライブラリのようなAnalyzerが存在する
		inspect.Analyzer,
	},
}

var pattenString = "^(is|has|are|ok|exists|exist|matched)"
var pattern = regexp.MustCompile(pattenString)

func run(pass *analysis.Pass) (any, error) {
	// inspect.Analyzerによる静的解析処理(依存する静的解析処理)の結果を取得する
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// (1) AST探索
	// Go言語で書かれたソースコードの中にあるbool型の変数を表すノードを探索する
	// ためにAST探索処理が実行される。
	// inspect.PreorderSeq関数は探索処理を実行する。
	// PreorderSeq関数の引数には探索時のフィルタを指定できる。
	// PreorderSeq関数は、ここに指定したノードを訪れた時に反復処理を実行する。
	for n := range inspect.PreorderSeq(
		(*ast.ValueSpec)(nil),
		(*ast.Field)(nil),
		(*ast.AssignStmt)(nil),
	) {
		// (2) 型チェック
		// pass.TypesInfoを用いて型チェックする
		// 型チェックすることで変数の型がブール値であるかどうかを判定できる
		//
		// pass.TypesInfo.TypeOf関数は、ast.Exprノードの型を返す
		// ast.Exprノードは、評価式を表す。
		// (例) ast.Exprが`1+2`なら、TypeOf関数の返り値はint型
		// (例) ast.Exprが`1+0.5`なら、TypeOf関数の返り値はfloat型
		// (例) ast.Exprが`a+b`なら、TypeOf関数の返り値はaとbに依存して型が決定される
		// go/typesパッケージには型決定アルゴリズムが実装されていて、それに従って型が決定されます
		switch n := n.(type) {
		case *ast.ValueSpec:
			// ast.ValueSpecノードは、varとかconst句による宣言文を表す
			for i, value := range n.Values {
				check(pass, n.Names[i], pass.TypesInfo.TypeOf(value))
			}
		case *ast.Field:
			// ast.Fieldノードは、関数、（構造体やインターフェースの）メソッドの引数、返り値のリストを表す
			if len(n.Names) <= 0 { // 名前がない引数や返り値である場合はスキップ
				continue
			}
			for _, name := range n.Names {
				check(pass, name, pass.TypesInfo.TypeOf(n.Type))
			}
		case *ast.AssignStmt:
			// ast.AssignStmtノードは、`:=`による値の割り当てを表す
			for _, lh := range n.Lhs {
				lIdent, ok := lh.(*ast.Ident)
				if !ok {
					panic(ok)
				}
				check(pass, lIdent, pass.TypesInfo.TypeOf(lh))
			}
		}
	}

	return nil, nil
}

func check(
	pass *analysis.Pass,
	paramName *ast.Ident,
	paramType types.Type,
) {
	if paramType == nil {
		return
	}

	// Underlying()関数は、その型のunderlying typeを取得します
	// underlying typeが初耳な人向け資料がこちら
	//   https://speakerdeck.com/dqneo/go-language-underlying-type
	// 型定義とか型エイリアスである場合があり得るため、underlying typeを取得する必要があります
	paramType = paramType.Underlying()

	if types.Identical(paramType, types.Typ[types.Bool]) ||
		types.Identical(paramType, types.Typ[types.UntypedBool]) {
		if matched := pattern.MatchString(paramName.Name); !matched {
			pass.Reportf(
				paramName.Pos(),
				"a boolean variable does not match pattern",
			)
		}
	}
}
