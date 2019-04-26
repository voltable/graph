package ast_test

import (
	"testing"

	"github.com/voltable/graph/query/cypher/ast"
)

func Test_MathPrecedence(t *testing.T) {
	var tests = []struct {
		expr   ast.MathematicalExpr
		result int
		err    string
	}{
		{
			expr:   &ast.AddExpr{},
			result: 2,
		},
		{
			expr:   &ast.SubtractExpr{},
			result: 2,
		},
		{
			expr:   &ast.DivideExpr{},
			result: 3,
		},
		{
			expr:   &ast.MultiplyExpr{},
			result: 3,
		},
		{
			expr:   &ast.ModuloExpr{},
			result: 4,
		},
		{
			expr:   &ast.PowerExpr{},
			result: 4,
		},
	}

	for i, tt := range tests {
		result := ast.MathPrecedence(tt.expr)
		if result != tt.result {
			t.Errorf("%d.  %q: comparison mismatch:\n  exp=%v\n  got=%v\n\n", i, tt.expr, tt.result, result)
		}
	}
}
