package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_ParenthesesPrecedence(t *testing.T) {

	p := ast.ParenthesesExpr{Parentheses: ast.LPAREN}
	if ast.ParenthesesPrecedence(p) != 11 {
		t.Errorf("LPAREN expected")
	}

	p.Parentheses = ast.RPAREN
	if ast.ParenthesesPrecedence(p) != 12 {
		t.Errorf("RPAREN expected")
	}

	p.Parentheses = 20
	if ast.ParenthesesPrecedence(p) != 20 {
		t.Errorf("Unknown expected")
	}
}
