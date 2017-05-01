package parser

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

type StackExpr []ast.Expr

func (s StackExpr) Push(v ast.Expr) StackExpr {
	return append(s, v)
}

func (s StackExpr) Pop() (StackExpr, ast.Expr, bool) {
	l := len(s)
	if l > 0 {
		return s[:l-1], s[l-1], true
		//return s[1:], s[0], true
	}
	return s, nil, false
}

// Shunt builds up the AST by Shunting the stack
func (s StackExpr) Shunt() (StackExpr, error) {
	var item ast.Expr
	var value ast.Expr

	valueStack := make(StackExpr, 0)
	comparisonStack := make(StackExpr, 0)
	comparisonCompletedStack := make(StackExpr, 0)
	booleanStack := make(StackExpr, 0)
	resultStack := make(StackExpr, 0)

	for len(s) > 0 {
		s, item, _ = s.Pop()

		// If the token is a value (value here includes both Ident and PropertyStmt).
		if _, ok := item.(*ast.Ident); ok {
			valueStack = valueStack.Push(item)
		} else if _, ok := item.(*ast.PropertyStmt); ok {
			valueStack = valueStack.Push(item)
		} else if _, ok := item.(*ast.ComparisonExpr); ok {
			// Otherwise, the token is an operator (operator here includes both ComparisonExpr and BooleanExpr).
			comparisonStack = comparisonStack.Push(item)
		} else {
			booleanStack = booleanStack.Push(item)

		}

		// If there are 2 values on the stack
		if len(valueStack) >= 2 {
			// Evaluate the operator, with the values as arguments.
			comparisonStack, item, _ = comparisonStack.Pop()
			if fun, ok := ast.IsOperatorWithFreeXorY(item); ok {
				valueStack, value, _ = valueStack.Pop()
				fun(value)
				if fun, ok := ast.IsOperatorWithFreeXorY(item); ok {
					valueStack, value, _ = valueStack.Pop()
					fun(value)
					//Push the returned results, if any, back onto the stack.
					comparisonCompletedStack = comparisonCompletedStack.Push(item)
				}
			}
		}

		// If there are 2 values on the stack
		if len(comparisonCompletedStack) >= 2 {
			// Evaluate the operator, with the values as arguments.
			booleanStack, item, _ = booleanStack.Pop()
			if fun, ok := ast.IsOperatorWithFreeXorY(item); ok {
				comparisonCompletedStack, value, _ = comparisonCompletedStack.Pop()
				fun(value)
				if fun, ok := ast.IsOperatorWithFreeXorY(item); ok {
					comparisonCompletedStack, value, _ = comparisonCompletedStack.Pop()
					fun(value)
					//Push the returned results, if any, back onto the stack.
					resultStack = resultStack.Push(item)
				}
			}
		}
	}

	return resultStack, nil
}
