package ast_test

import (
	"math"
	"testing"

	"github.com/RossMerr/Caudex.Graph/comparisons"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_Interpret(t *testing.T) {

	var tests = []struct {
		c      ast.ComparisonExpr
		v      *vertices.Vertex
		p      *ast.VertexPatn
		result bool
		err    string
	}{
		{
			c:      ast.ComparisonExpr{Comparison: comparisons.EQ},
			v:      &vertices.Vertex{},
			p:      &ast.VertexPatn{},
			result: true,
		},
		{
			c:      ast.ComparisonExpr{Comparison: comparisons.NEQ, Left: ast.PropertyStmt{}, Right: ast.Ident{}},
			v:      &vertices.Vertex{},
			p:      &ast.VertexPatn{Variable: "Person"},
			result: true,
		},
		{
			c:      ast.ComparisonExpr{Comparison: comparisons.IS_NULL, Left: ast.PropertyStmt{Variable: "n"}, Right: ast.Ident{}},
			v:      &vertices.Vertex{},
			p:      &ast.VertexPatn{Variable: "n"},
			result: true,
		},
		{
			c: ast.ComparisonExpr{Comparison: comparisons.IS_NOT_NULL, Left: ast.PropertyStmt{Variable: "n", Value: "Person"}, Right: ast.Ident{}},
			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("Person", "John Smith")
				return x
			}(),
			p:      &ast.VertexPatn{Variable: "n"},
			result: true,
		},
		{
			c: ast.ComparisonExpr{Comparison: comparisons.LT, Left: ast.PropertyStmt{Variable: "n", Value: "Age"}, Right: ast.Ident{Data: math.MaxInt32}},
			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("Age", math.MaxInt32-1)
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

func Test_ComparisonPrecedence(t *testing.T) {

	var tests = []struct {
		c ast.ComparisonExpr
		p int
	}{
		{
			c: ast.ComparisonExpr{Comparison: comparisons.EQ},
			p: 8,
		}, {
			c: ast.ComparisonExpr{Comparison: comparisons.NEQ},
			p: 8,
		}, {
			c: ast.ComparisonExpr{Comparison: comparisons.LT},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: comparisons.LTE},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: comparisons.GT},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: comparisons.GTE},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: comparisons.IS_NULL},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: comparisons.IS_NOT_NULL},
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

// Swaps the boolean
func XOR(b bool) bool {
	if b == true {
		return false
	}

	return true
}
