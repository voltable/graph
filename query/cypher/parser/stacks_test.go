package parser_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
)

// n.number >= 1 AND n.number <= 10
func TestBasic_UpdateStack(t *testing.T) {
	exprStack := make(parser.StackExpr, 0)

	n1 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.Push(n1)

	n2 := &ast.ComparisonExpr{Comparison: ast.GTE}
	exprStack = exprStack.Push(n2)

	n3 := &ast.Ident{Data: 1}
	exprStack = exprStack.Push(n3)

	n4 := &ast.BooleanExpr{Boolean: ast.AND}
	exprStack = exprStack.Push(n4)

	n5 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.Push(n5)

	n6 := &ast.ComparisonExpr{Comparison: ast.NEQ}
	exprStack = exprStack.Push(n6)

	n7 := &ast.Ident{Data: 10}
	exprStack = exprStack.Push(n7)

	result, _ := exprStack.Shunt()

	if result != n4 {
		t.Errorf("found %s expected %s", result, n4)
	}

	if n4.X != n2 {
		t.Errorf("found %s expected %s", n4.X, n2)
	}

	if n2.X != n1 {
		t.Errorf("found %s expected %s", n2.X, n1)
	}

	if n2.Y != n3 {
		t.Errorf("found %s expected %s", n2.Y, n3)
	}

	if n4.Y != n6 {
		t.Errorf("found %s expected %s", n4.Y, n6)
	}

	if n6.X != n5 {
		t.Errorf("found %s expected %s", n6.X, n5)
	}

	if n6.Y != n7 {
		t.Errorf("found %s expected %s", n6.Y, n7)
	}
}

// n.name = 'Peter' XOR (n.age < 30 AND n.name = 'Tobias') OR NOT (n.name = 'Tobias' OR n.name = 'Peter')
func TestDeep_UpdateStack(t *testing.T) {
	exprStack := make(parser.StackExpr, 0)

	n1 := &ast.PropertyStmt{Variable: "n", Value: "name"}
	exprStack = exprStack.Push(n1)

	n2 := &ast.ComparisonExpr{Comparison: ast.EQ}
	exprStack = exprStack.Push(n2)

	n3 := &ast.Ident{Data: "Peter"}
	exprStack = exprStack.Push(n3)

	n4 := &ast.BooleanExpr{Boolean: ast.AND}
	exprStack = exprStack.Push(n4)

	n5 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.Push(n5)

	n6 := &ast.ComparisonExpr{Comparison: ast.NEQ}
	exprStack = exprStack.Push(n6)

	n7 := &ast.Ident{Data: 10}
	exprStack = exprStack.Push(n7)

	exprStack.Shunt()
}
