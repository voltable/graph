package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_Evaluate(t *testing.T) {
	var tests = []struct {
		c      ast.ComparisonExpr
		v      *vertices.Vertex
		p      *ast.VertexPatn
		result bool
		err    string
	}{
		{
			c:      ast.ComparisonExpr{Comparison: ast.EQ},
			v:      &vertices.Vertex{},
			p:      &ast.VertexPatn{},
			result: false,
		},
		{
			c:      ast.ComparisonExpr{Comparison: ast.EQ},
			v:      &vertices.Vertex{},
			p:      &ast.VertexPatn{},
			result: false,
		},
	}

	for _, tt := range tests {
		tt.c.Evaluate(tt.v, tt.p)
	}
}

func Test_ComparisonPrecedence(t *testing.T) {

	var tests = []struct {
		c ast.ComparisonExpr
		p int
	}{
		{
			c: ast.ComparisonExpr{Comparison: ast.EQ},
			p: 8,
		}, {
			c: ast.ComparisonExpr{Comparison: ast.NEQ},
			p: 8,
		}, {
			c: ast.ComparisonExpr{Comparison: ast.LT},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: ast.LTE},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: ast.GT},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: ast.GTE},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: ast.IS_NULL},
			p: 7,
		}, {
			c: ast.ComparisonExpr{Comparison: ast.IS_NOT_NULL},
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
