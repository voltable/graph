package ast_test

import (
	"fmt"
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
		// {
		// 	c:      ast.ComparisonExpr{Comparison: ast.EQ},
		// 	v:      &vertices.Vertex{},
		// 	p:      &ast.VertexPatn{},
		// 	result: true,
		// },
		// {
		// 	c:      ast.ComparisonExpr{Comparison: ast.NEQ, X: ast.PropertyStmt{}, Y: ast.Ident{}},
		// 	v:      &vertices.Vertex{},
		// 	p:      &ast.VertexPatn{Variable: "Person"},
		// 	result: true,
		// },
		// {
		// 	c:      ast.ComparisonExpr{Comparison: ast.IS_NULL, X: ast.PropertyStmt{Variable: "n"}, Y: ast.Ident{}},
		// 	v:      &vertices.Vertex{},
		// 	p:      &ast.VertexPatn{Variable: "n"},
		// 	result: true,
		// },
		// {
		// 	c: ast.ComparisonExpr{Comparison: ast.IS_NOT_NULL, X: ast.PropertyStmt{Variable: "n", Value: "Person"}, Y: ast.Ident{}},
		// 	v: func() *vertices.Vertex {
		// 		x, _ := vertices.NewVertex()
		// 		x.SetProperty("Person", "John Smith")
		// 		return x
		// 	}(),
		// 	p:      &ast.VertexPatn{Variable: "n"},
		// 	result: true,
		// },
		{
			c: ast.ComparisonExpr{Comparison: ast.LT, X: ast.PropertyStmt{Variable: "n", Value: "Age"}, Y: ast.Ident{Data: 1}},
			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()

				x.SetPropertyInt("Age", 10)
				return x
			}(),
			p:      &ast.VertexPatn{Variable: "n"},
			result: true,
		},
	}
	x, _ := vertices.NewVertex()

	x.SetPropertyInt("Age", 10)
	fmt.Printf("age %q\n", x.Property("Age"))

	for i, tt := range tests {
		result := tt.c.Evaluate(tt.v, tt.p)
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
