package parser_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/query/cypher/parser"
)

// n.number >= 1 AND n.number <= 10
func Test_UpdateStack(t *testing.T) {
	exprStack := make(parser.StackExpr, 0)

	n1 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.UpdateStack(n1)

	n2 := &ast.ComparisonExpr{Comparison: ast.GTE}
	exprStack = exprStack.UpdateStack(n2)

	n3 := &ast.Ident{Data: 1}
	exprStack = exprStack.UpdateStack(n3)

	n4 := &ast.BooleanExpr{Boolean: ast.AND}
	exprStack = exprStack.UpdateStack(n4)

	n5 := &ast.PropertyStmt{Variable: "n", Value: "number"}
	exprStack = exprStack.UpdateStack(n5)

	n6 := &ast.ComparisonExpr{Comparison: ast.NEQ}
	exprStack = exprStack.UpdateStack(n6)

	n7 := &ast.Ident{Data: 10}
	exprStack = exprStack.UpdateStack(n7)
}
