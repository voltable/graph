package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/keyvalue"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_BooleanPrecedence(t *testing.T) {

	var tests = []struct {
		c        ast.BooleanExpr
		p        int
		expected expressions.Boolean
		err      string
	}{
		{
			c:        ast.BooleanExpr{Boolean: expressions.AND},
			p:        9,
			expected: expressions.AND,
		},
		{
			c:        ast.BooleanExpr{Boolean: expressions.OR},
			p:        11,
			expected: expressions.OR,
		},
		{
			c:        ast.BooleanExpr{Boolean: expressions.XOR},
			p:        10,
			expected: expressions.XOR,
		},
		{
			c:        ast.BooleanExpr{Boolean: 100},
			p:        20,
			expected: 20,
		},
	}

	for i, tt := range tests {
		if ast.BooleanPrecedence(tt.c) != tt.p {
			t.Errorf("%d. boolean expected %v", i, tt.expected)
		}
	}

}

func Test_BooleanGetSet(t *testing.T) {
	c := ast.BooleanExpr{Boolean: expressions.AND}
	c.SetLeft(nil)
	c.SetRight(nil)
	if c.GetLeft() != nil {
		t.Errorf("boolean left not nil got:\n got=%t\n\n", c.GetLeft())
	}
	if c.GetRight() != nil {
		t.Errorf("boolean right not nil got :\n got=%t\n\n", c.GetRight())
	}
}

func Test_BooleanExprInterpret(t *testing.T) {

	var tests = []struct {
		c      *ast.BooleanExpr
		v      *keyvalue.KeyValue
		p      string
		result bool
		err    string
	}{
		{
			c: ast.NewBooleanExpr(expressions.AND,
				ast.NewComparisonExpr(expressions.GT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 10}),
				ast.NewComparisonExpr(expressions.LT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 1000}),
			),
			v: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue(100, id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("Age"))
				return x
			}(),
			p:      "n",
			result: true,
		},
		{
			c: ast.NewBooleanExpr(expressions.AND,
				ast.NewComparisonExpr(expressions.LT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 10}),
				ast.NewComparisonExpr(expressions.GT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 1000}),
			),
			v: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue(100, id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("Age"))
				return x
			}(),
			p:      "n",
			result: false,
		},
		{
			c: ast.NewBooleanExpr(expressions.OR,
				ast.NewComparisonExpr(expressions.GT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 10}),
				ast.NewComparisonExpr(expressions.LT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 1000}),
			),
			v: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue(100, id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("Age"))
				return x
			}(),
			p:      "n",
			result: true,
		},
		{
			c: ast.NewBooleanExpr(expressions.OR,
				ast.NewComparisonExpr(expressions.LT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 10}),
				ast.NewComparisonExpr(expressions.GT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 1000}),
			),
			v: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue(100, id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("Age"))
				return x
			}(),
			p:      "n",
			result: false,
		},
		{
			c: ast.NewBooleanExpr(expressions.XOR,
				ast.NewComparisonExpr(expressions.LT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 10}),
				ast.NewComparisonExpr(expressions.LT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 1000}),
			),
			v: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue(100, id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("Age"))
				return x
			}(),
			p:      "n",
			result: true,
		},
		{
			c: ast.NewBooleanExpr(expressions.XOR,
				ast.NewComparisonExpr(expressions.LT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 10}),
				ast.NewComparisonExpr(expressions.GT, &ast.PropertyStmt{Variable: "n", Value: "Age"}, &ast.Ident{Data: 1000}),
			),
			v: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue(100, id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("Age"))
				return x
			}(),
			p:      "n",
			result: false,
		},
		{
			c: ast.NewBooleanExpr(20, nil, nil),
			v: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue(100, id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("Age"))
				return x
			}(),
			p:      "n",
			result: false,
		},
	}

	for i, tt := range tests {
		result := tt.c.Interpret(tt.p, tt.v)
		if result != tt.result {
			t.Errorf("%d.  %q: comparison mismatch:\n  exp=%t\n  got=%t\n\n", i, tt.c, tt.result, result)
		}
	}
}
