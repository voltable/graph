package cypher_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/keyvalue"
	"github.com/RossMerr/Caudex.Graph/uuid"

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
					id, _ := uuid.GenerateRandomUUID()
					v := keyvalue.NewKeyValue("foo", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))

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
					id, _ := uuid.GenerateRandomUUID()
					v := keyvalue.NewKeyValue("foo", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))

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
					x := &keyvalue.KeyValue{}
					v := &keyvalue.KeyValue{}

					f := query.NewFrontier(x, "")
					fq := f.Values[0]
					fv := query.FrontierProperties{KeyValue: v, Variable: ""}
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
					x := &keyvalue.KeyValue{}
					v := &keyvalue.KeyValue{}

					to, _ := uuid.GenerateRandomUUID()
					from, _ := uuid.GenerateRandomUUID()

					e := keyvalue.NewKeyValue(string(to[:]), from[:], keyvalue.US, keyvalue.Relationship, keyvalue.US, []byte(""))

					f := query.NewFrontier(x, "")
					fq := f.Values[0]
					fv := query.FrontierProperties{KeyValue: v, Variable: ""}
					fe := query.FrontierProperties{KeyValue: e, Variable: ""}

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
		v        *keyvalue.KeyValue
		result   bool
	}{
		// 0
		{
			expr:     &ast.PropertyStmt{Variable: "n", Value: "name"},
			variable: "n",
			v: func() *keyvalue.KeyValue {
				id, _ := uuid.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("foo", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
			}(),
			result: false,
		},
		// 1
		{
			expr:     ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}),
			variable: "n",
			v: func() *keyvalue.KeyValue {
				id, _ := uuid.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("foo", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
			}(),
			result: true,
		},
		// 2
		{
			expr:     ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}),
			variable: "x",
			v: func() *keyvalue.KeyValue {
				id, _ := uuid.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("foo", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
			}(),
			result: false,
		},
		// 3
		{
			expr:     ast.NewBooleanExpr(expressions.OR, ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}), nil),
			variable: "n",
			v: func() *keyvalue.KeyValue {
				id, _ := uuid.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("foo", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
			}(),
			result: true,
		},
		// 4
		{
			expr: ast.NewBooleanExpr(expressions.OR,
				ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}),
				ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "m", Value: "name"}, &ast.Ident{Data: "bar"}),
			),
			variable: "n",
			v: func() *keyvalue.KeyValue {
				id, _ := uuid.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("foo", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
			}(),
			result: true,
		},
		// 5
		{
			expr:     ast.NewBooleanExpr(expressions.OR, ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}), ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "m", Value: "name"}, &ast.Ident{Data: "bar"})),
			variable: "m",
			v: func() *keyvalue.KeyValue {
				id, _ := uuid.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("foo", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
			}(),
			result: false,
		},
		// 6
		{
			expr: ast.NewBooleanExpr(expressions.OR,
				ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}),
				ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "m", Value: "name"}, &ast.Ident{Data: "bar"}),
			),
			variable: "m",
			v: func() *keyvalue.KeyValue {
				id, _ := uuid.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("bar", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
			}(),
			result: true,
		},
		// 7
		{
			expr:     ast.NewBooleanExpr(expressions.OR, ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "n", Value: "name"}, &ast.Ident{Data: "foo"}), ast.NewComparisonExpr(expressions.EQ, &ast.PropertyStmt{Variable: "m", Value: "person"}, &ast.Ident{Data: "john smith"})),
			variable: "m",
			v: func() *keyvalue.KeyValue {
				id, _ := uuid.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("john smith", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("person"))
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
