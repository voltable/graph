package cypher_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/vertices"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_Filter(t *testing.T) {
	filter := cypher.NewFilter()
	var tests = []struct {
		predicate ast.Expr
		expected  int
		iterate   int
		setup     func(int) query.IteratorFrontier
	}{
		// 0
		{
			setup: func(iterate int) query.IteratorFrontier {
				return func() (*query.Frontier, bool) {
					return nil, false
				}
			},
			predicate: nil,
			iterate:   1,
			expected:  0,
		},
		// 1
		{
			setup: func(iterate int) query.IteratorFrontier {
				count := 0
				return func() (*query.Frontier, bool) {
					f := query.Frontier{}
					if count < iterate {
						count++
						return &f, true
					}
					return &f, false
				}
			},
			iterate:   1,
			predicate: ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}),
			expected:  0,
		},
		// 2
		{
			setup: func(iterate int) query.IteratorFrontier {
				count := 0
				return func() (*query.Frontier, bool) {
					v, _ := vertices.NewVertex()
					v.SetProperty("name", "foo")
					f := query.NewFrontier(v, "n")
					if count < iterate {
						count++
						return &f, true
					}
					return &f, false
				}
			},
			iterate:   1,
			predicate: ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}),
			expected:  1,
		},
		// 3
		{
			setup: func(iterate int) query.IteratorFrontier {
				count := 0
				return func() (*query.Frontier, bool) {
					v, _ := vertices.NewVertex()
					v.SetProperty("name", "foo")
					f := query.NewFrontier(v, "n")
					if count < iterate {
						count++
						return &f, true
					}
					return &f, false
				}
			},
			iterate:   1,
			predicate: nil,
			expected:  1,
		},
		// 4
		{
			setup: func(iterate int) query.IteratorFrontier {
				count := 0
				return func() (*query.Frontier, bool) {
					x, _ := vertices.NewVertex()
					v, _ := vertices.NewVertex()
					f := query.NewFrontier(x, "")
					fq := f.Values[0]
					fv := &query.FrontierVertex{Vertex: v, Variable: ""}
					fq.Parts = append(fq.Parts, fv)

					if count < iterate {
						count++
						return &f, true
					}
					return &f, false
				}
			},
			iterate:   1,
			predicate: nil,
			expected:  2,
		},
		// 5
		{
			predicate: nil,
			expected:  3,
			iterate:   1,
			setup: func(iterate int) query.IteratorFrontier {
				count := 0
				return func() (*query.Frontier, bool) {
					x, _ := vertices.NewVertex()
					v, _ := vertices.NewVertex()
					e, _ := x.AddDirectedEdge(v)
					f := query.NewFrontier(x, "")
					fq := f.Values[0]
					fv := &query.FrontierVertex{Vertex: v, Variable: ""}
					fe := &query.FrontierEdge{Edge: e, Variable: ""}

					fq.Parts = append(fq.Parts, fe)
					fq.Parts = append(fq.Parts, fv)

					if count < iterate {
						count++
						return &f, true
					}
					return &f, false
				}
			},
		},
	}

	for i, tt := range tests {
		result := filter.Filter(tt.setup(tt.iterate), tt.predicate)
		count := 0
		for v, ok := result(); ok; v, ok = result() {
			count += len(v.OptimalPath())

			if v == nil {
				t.Errorf("%d %+v", i, v)
			}
		}
		if count != tt.expected {
			t.Errorf("%d. expected %d got %d", i, tt.expected, count)
		}
	}
}

func Test_ExpressionEvaluator(t *testing.T) {

	filter := cypher.NewFilter()
	var tests = []struct {
		expr     ast.Expr
		variable string
		v        *vertices.Vertex
		result   bool
	}{
		{
			expr:     &ast.PropertyStmt{Variable: "n", Value: "name"},
			variable: "n",
			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("name", "foo")
				return x
			}(),
			result: false,
		},
		{
			expr:     ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}),
			variable: "n",
			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("name", "foo")
				return x
			}(),
			result: true,
		},
		{
			expr:     ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}),
			variable: "x",
			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("name", "foo")
				return x
			}(),
			result: false,
		},
		{
			expr:     ast.NewBooleanExpr(expressions.OR, ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}), nil),
			variable: "n",
			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("name", "foo")
				return x
			}(),
			result: true,
		},
		{
			expr:     ast.NewBooleanExpr(expressions.OR, ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}), ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "m", Value: "name"}, &ast.Ident{Data: "bar"})),
			variable: "n",
			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("name", "foo")
				return x
			}(),
			result: true,
		},
		{
			expr:     ast.NewBooleanExpr(expressions.OR, ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}), ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "m", Value: "name"}, &ast.Ident{Data: "bar"})),
			variable: "m",
			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("name", "foo")
				return x
			}(),
			result: false,
		},
		{
			expr:     ast.NewBooleanExpr(expressions.OR, ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}), ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "m", Value: "name"}, &ast.Ident{Data: "bar"})),
			variable: "m",
			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("name", "bar")
				return x
			}(),
			result: true,
		},
		{
			expr:     ast.NewBooleanExpr(expressions.OR, ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}), ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "m", Value: "person"}, &ast.Ident{Data: "john smith"})),
			variable: "m",
			v: func() *vertices.Vertex {
				x, _ := vertices.NewVertex()
				x.SetProperty("person", "john smith")
				return x
			}(),
			result: true,
		},
	}
	for i, tt := range tests {
		result := filter.ExpressionEvaluator(tt.expr, tt.variable, tt.v)
		if result != tt.result {
			t.Errorf("%d. exp:\n %+v got:\n %+v", i, tt.result, result)
		}
	}
}
