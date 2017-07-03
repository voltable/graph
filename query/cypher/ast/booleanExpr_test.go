package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_BooleanPrecedence(t *testing.T) {
	c := ast.BooleanExpr{Boolean: ast.AND}

	if ast.BooleanPrecedence(c) != 9 {
		t.Errorf("boolean expected %v", ast.AND)
	}

	c.Boolean = ast.OR
	if ast.BooleanPrecedence(c) != 11 {
		t.Errorf("boolean expected %v", ast.OR)
	}

	c.Boolean = ast.XOR
	if ast.BooleanPrecedence(c) != 10 {
		t.Errorf("boolean expected %v", ast.XOR)
	}

	c.Boolean = 100
	if ast.BooleanPrecedence(c) != 20 {
		t.Errorf("boolean expected")
	}
}

func Test_NotPrecedence(t *testing.T) {
	c := ast.NotExpr{}

	if ast.NotPrecedence(c) != 13 {
		t.Errorf("not expected")
	}
}
