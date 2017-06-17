package parser

import (
	"fmt"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

type StackExpr []ast.Expr

// Push add's a item the the StackExpr
func (s StackExpr) Push(v ast.Expr) StackExpr {
	return append(s, v)
}

// pop removes the last item on the StackExpr and returns it
func (s StackExpr) pop() (StackExpr, ast.Expr) {
	l := len(s)
	if l > 0 {
		return s[:l-1], s[l-1]
	}
	return s, nil
}

// pop removes the first item on the StackExpr and returns it
func (s StackExpr) shift() (StackExpr, ast.Expr) {
	l := len(s)
	if l > 0 {
		return s[1:], s[0]
	}
	return s, nil
}

// top returns the last item on the StackExpr without removing it
func (s StackExpr) top() ast.Expr {
	l := len(s)
	if l > 0 {
		return s[l-1]
	}
	return nil
}

// Shunt builds up the AST by Shunting the stack
func (s StackExpr) Shunt() (ast.Expr, error) {
	var item ast.Expr

	exprStack := make(StackExpr, 0)
	operatorStack := make(StackExpr, 0)
	notStack := make(StackExpr, 0)

	for len(s) > 0 {
		s, item = s.shift()
		if p, ok := item.(*ast.ParenthesesExpr); ok {
			if p.Parentheses == ast.LPAREN {
				operatorStack = operatorStack.Push(item)
			} else { // RPAREN
				var expr ast.Expr
				operatorStack, expr = operatorStack.pop()
				for expr != nil {
					if p, ok := expr.(*ast.ParenthesesExpr); ok && p.Parentheses == ast.LPAREN {
						break
					} else {
						operatorStack, exprStack, notStack = buildExprFromOperator(expr, operatorStack, exprStack, notStack)
					}
					operatorStack, expr = operatorStack.pop()

				}
			}
		} else if _, ok := item.(*ast.Ident); ok {
			// If the token is a value (value here includes both Ident and PropertyStmt).
			fmt.Printf("%s went on exprStack \n", item)
			exprStack = exprStack.Push(item)
		} else if _, ok := item.(*ast.PropertyStmt); ok {
			// If the token is a value (value here includes both Ident and PropertyStmt).
			fmt.Printf("%s went on exprStack \n", item)
			exprStack = exprStack.Push(item)
		} else if _, ok := item.(*ast.ComparisonExpr); ok {
			// Otherwise, the token is an operator (operator here includes both ComparisonExpr and BooleanExpr).
			operatorStack, exprStack, notStack = shuntOperator(item, operatorStack, exprStack, notStack)
		} else if b, ok := item.(*ast.BooleanExpr); ok {
			if b.Boolean == ast.NOT {
				notStack = notStack.Push(item)
			} else {
				// Otherwise, the token is an operator (operator here includes both ComparisonExpr and BooleanExpr).
				operatorStack, exprStack, notStack = shuntOperator(item, operatorStack, exprStack, notStack)
			}
		}
	}

	// while there are still operators on the operatorStack:
	for len(operatorStack) > 0 {
		var expr ast.Expr
		operatorStack, expr = operatorStack.pop()
		operatorStack, exprStack, notStack = buildExprFromOperator(expr, operatorStack, exprStack, notStack)
	}

	var result ast.Expr

	exprStack, result = exprStack.pop()
	return result, nil
}

func shuntOperator(item ast.Expr, operatorStack StackExpr, exprStack StackExpr, notStack StackExpr) (StackExpr, StackExpr, StackExpr) {

	for expr := operatorStack.top(); expr != nil && ast.Precedence(expr) <= ast.Precedence(item); expr = operatorStack.top() {
		operatorStack, expr = operatorStack.pop()
		operatorStack, exprStack, notStack = buildExprFromOperator(expr, operatorStack, exprStack, notStack)
	}

	operatorStack = operatorStack.Push(item)

	return operatorStack, exprStack, notStack
}

func buildExprFromOperator(expr ast.Expr, operatorStack StackExpr, exprStack StackExpr, notStack StackExpr) (StackExpr, StackExpr, StackExpr) {
	var y, x ast.Expr
	if operator, ok := expr.(ast.OperatorExpr); ok {
		exprStack, y = exprStack.pop()
		exprStack, x = exprStack.pop()

		operator.SetY(y)
		operator.SetX(x)

		// If we find anything on the notStack we should make the operator a child of it
		if len(notStack) > 0 {
			var n ast.Expr
			notStack, n = notStack.pop()
			if not, ok := n.(ast.OperatorExpr); ok {
				not.SetX(expr)
				exprStack = exprStack.Push(n)
			}
		} else {
			exprStack = exprStack.Push(expr)
		}
	}

	return operatorStack, exprStack, notStack
}
