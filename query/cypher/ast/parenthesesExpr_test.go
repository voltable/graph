package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_ParenthesesPrecedence(t *testing.T) {
	var tests = []struct {
		expr   ast.ParenthesesExpr
		result int
		err    string
	}{
		{
			expr:   ast.ParenthesesExpr{Parentheses: ast.LPAREN},
			result: 11,
		},
		{
			expr:   ast.ParenthesesExpr{Parentheses: ast.RPAREN},
			result: 12,
		},
		{
			expr:   ast.ParenthesesExpr{Parentheses: 20},
			result: 20,
		},
	}

	for i, tt := range tests {
		result := ast.ParenthesesPrecedence(tt.expr)
		if result != tt.result {
			t.Errorf("%d.  %q: comparison mismatch:\n  exp=%v\n  got=%v\n\n", i, tt.expr, tt.result, result)
		}
	}
}
