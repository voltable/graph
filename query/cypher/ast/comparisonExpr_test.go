package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_ComparisonPrecedence(t *testing.T) {

	c := ast.ComparisonExpr{Comparison: ast.EQ}

	if ast.ComparisonPrecedence(c) != 8 {
		t.Errorf("comparison expected %v", ast.EQ)
	}

	c.Comparison = ast.NEQ
	if ast.ComparisonPrecedence(c) != 8 {
		t.Errorf("comparison expected %v", ast.NEQ)
	}

	c.Comparison = ast.LT
	if ast.ComparisonPrecedence(c) != 7 {
		t.Errorf("comparison expected %v", ast.LT)
	}

	c.Comparison = ast.LTE
	if ast.ComparisonPrecedence(c) != 7 {
		t.Errorf("comparison expected %v", ast.LTE)
	}

	c.Comparison = ast.GT
	if ast.ComparisonPrecedence(c) != 7 {
		t.Errorf("comparison expected %v", ast.GT)
	}

	c.Comparison = ast.GTE
	if ast.ComparisonPrecedence(c) != 7 {
		t.Errorf("comparison expected %v", ast.GTE)
	}

	c.Comparison = ast.IS
	if ast.ComparisonPrecedence(c) != 7 {
		t.Errorf("comparison expected %v", ast.IS)
	}

	c.Comparison = ast.NULL
	if ast.ComparisonPrecedence(c) != 7 {
		t.Errorf("comparison expected %v", ast.NULL)
	}

	c.Comparison = 10
	if ast.ComparisonPrecedence(c) != 20 {
		t.Errorf("comparison expected")
	}
}
