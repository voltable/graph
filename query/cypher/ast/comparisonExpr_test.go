package ast_test

import (
	"math"
	"testing"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_ComparisonExprInterpret(t *testing.T) {

	var tests = []struct {
		c      *ast.ComparisonExpr
		v      *graph.Vertex
		p      string
		result bool
		err    string
	}{
		{
			c:      &ast.ComparisonExpr{Comparison: expressions.EQ},
			v:      &graph.Vertex{},
			result: true,
		},
		{
			c:      ast.NewComparisonExpr(expressions.NEQ, &ast.PropertyStmt{}, &ast.Ident{}),
			v:      &graph.Vertex{},
			result: false,
		},
		{
			c:      ast.NewComparisonExpr(expressions.IS_NULL, &ast.PropertyStmt{Variable: "n"}, &ast.Ident{}),
			v:      &graph.Vertex{},
			p:      "n",
			result: true,
		},
		{
			c: ast.NewComparisonExpr(expressions.IS_NOT_NULL, &ast.PropertyStmt{Variable: "n", Value: "Person"}, &ast.Ident{}),
			v: func() *graph.Vertex {
				x, _ := graph.NewVertex()
				x.SetProperty("Person", "John Smith")
				return x
			}(),
			p:      "n",
			result: true,
		},
		{
			c: ast.NewComparisonExpr(expressions.LT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: math.MaxInt32}),
			v: func() *graph.Vertex {
				x, _ := graph.NewVertex()
				x.SetProperty("Age", math.MaxInt32-1)
				return x
			}(),
			p: "n",

			result: true,
		},
	}

	for i, tt := range tests {
		result := tt.c.Interpret(tt.p, tt.v)
		if result != tt.result {
			t.Errorf("%d. %q: comparison mismatch:\n  exp=%t\n  got=%t\n\n", i, tt.c, tt.result, result)
		}
	}
}

func Test_ComparisonPrecedence(t *testing.T) {

	var tests = []struct {
		c ast.ComparisonExpr
		p int
	}{
		{
			c: ast.ComparisonExpr{Comparison: expressions.EQ},
			p: 8,
		}, {
			c: ast.ComparisonExpr{Comparison: expressions.NEQ},
			p: 8,
		}, {
			c: ast.ComparisonExpr{Comparison: expressions.LT},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: expressions.LTE},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: expressions.GT},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: expressions.GTE},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: expressions.IS_NULL},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: expressions.IS_NOT_NULL},
			p: 7,
		},
	}

	for i, tt := range tests {
		precedence := ast.ComparisonPrecedence(tt.c)
		if precedence != tt.p {
			t.Errorf("%d. %q: comparison mismatch:\n  exp=%d\n  got=%d\n\n", i, tt.c, tt.p, precedence)
		}
	}
}
