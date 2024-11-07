package linter2

import (
	"go/ast"
	"go/types"
	"regexp"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = `linter2 is boolean variable naming checker`

var pattenString = "^(is|has|are)"
var pattern = regexp.MustCompile(pattenString)

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "linter2",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

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
