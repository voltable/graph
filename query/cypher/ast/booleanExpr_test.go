package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_BooleanPrecedence(t *testing.T) {
	c := ast.BooleanExpr{Boolean: expressions.AND}

	if ast.BooleanPrecedence(c) != 9 {
		t.Errorf("boolean expected %v", expressions.AND)
	}

	c.Boolean = expressions.OR
	if ast.BooleanPrecedence(c) != 11 {
		t.Errorf("boolean expected %v", expressions.OR)
	}

	c.Boolean = expressions.XOR
	if ast.BooleanPrecedence(c) != 10 {
		t.Errorf("boolean expected %v", expressions.XOR)
	}

	c.Boolean = 100
	if ast.BooleanPrecedence(c) != 20 {
		t.Errorf("boolean expected")
	}
}

func Test_BooleanExprInterpret(t *testing.T) {

	var tests = []struct {
		c      *ast.BooleanExpr
		v      *vertices.Vertex
		p      *ast.VertexPatn
		result bool
		err    string
	}{
		{
			c: ast.NewBooleanExpr(expressions.AND,
				ast.NewComparisonExpr(expressions.GT, ast.PropertyStmt{Variable: "n", Value: "Age"}, ast.Ident{Data: 10}),
				ast.NewComparisonExpr(expressions.LT, ast.PropertyStmt{Variable: "n", Value: "Age"}, ast.Ident{Data: 1000}),
			),

			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("Age", 100)
				return x
			}(),
			p:      &ast.VertexPatn{Variable: "n"},
			result: true,
		},
		{
			c: ast.NewBooleanExpr(expressions.OR,
				ast.NewComparisonExpr(expressions.GT, ast.PropertyStmt{Variable: "n", Value: "Age"}, ast.Ident{Data: 10}),
				ast.NewComparisonExpr(expressions.LT, ast.PropertyStmt{Variable: "n", Value: "Age"}, ast.Ident{Data: 1000}),
			),

			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("Age", 100)
				return x
			}(),
			p:      &ast.VertexPatn{Variable: "n"},
			result: true,
		},
		{
			c: ast.NewBooleanExpr(expressions.XOR,
				ast.NewComparisonExpr(expressions.LT, ast.PropertyStmt{Variable: "n", Value: "Age"}, ast.Ident{Data: 10}),
				ast.NewComparisonExpr(expressions.LT, ast.PropertyStmt{Variable: "n", Value: "Age"}, ast.Ident{Data: 1000}),
			),

			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("Age", 100)
				return x
			}(),
			p:      &ast.VertexPatn{Variable: "n"},
			result: true,
		},
	}

	for i, tt := range tests {
		result := tt.c.Interpret(tt.v, tt.p)
		if result != tt.result {
			t.Errorf("%d.  %q: comparison mismatch:\n  exp=%t\n  got=%t\n\n", i, tt.c, tt.result, result)
		}
	}
}
