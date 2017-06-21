package parser

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

// StackExpr a simple stack for the AST
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
	var item, result ast.Expr
	exprStack := make(StackExpr, 0)
	operatorStack := make(StackExpr, 0)
	for len(s) > 0 {
		s, item = s.shift()
		if p, ok := item.(*ast.ParenthesesExpr); ok {
			if p.Parentheses == ast.LPAREN {
				operatorStack = operatorStack.Push(item)
			} else { // RPAREN
				for expr := operatorStack.top(); expr != nil; expr = operatorStack.top() {
					if p, ok := expr.(*ast.ParenthesesExpr); ok && p.Parentheses == ast.LPAREN {
						break
					} else {
						operatorStack, exprStack, _ = shuntOperator(expr, operatorStack, exprStack)
					}
				}
				// remove the LPAREN
				operatorStack, _ = operatorStack.pop()
			}
		} else if ok := isValue(item); ok {
			// If the token is a value (value here includes both Ident and PropertyStmt).
			exprStack = exprStack.Push(item)
		} else if ok := isOperator(item); ok {
			// Otherwise, the token is an operator (operator here includes both ComparisonExpr and BooleanExpr).
			for expr := operatorStack.top(); expr != nil && ast.Precedence(expr) <= ast.Precedence(item); expr = operatorStack.top() {
				if operatorStack, exprStack, ok = shuntOperator(expr, operatorStack, exprStack); !ok {
					break
				}
			}
			operatorStack = operatorStack.Push(item)
		}
	}

	// while there are still operators on the operatorStack:
	for len(operatorStack) > 0 {
		expr := operatorStack.top()
		operatorStack, exprStack, _ = shuntOperator(expr, operatorStack, exprStack)
	}

	exprStack, result = exprStack.pop()
	return result, nil
}

func isValue(item ast.Expr) bool {
	if _, ok := item.(*ast.Ident); ok {
		return true
	} else if _, ok := item.(*ast.PropertyStmt); ok {
		return true
	}
	return false
}

func isOperator(item ast.Expr) bool {
	if _, ok := item.(*ast.ComparisonExpr); ok {
		return true
	} else if _, ok := item.(*ast.BooleanExpr); ok {
		return true
	} else if _, ok := item.(*ast.NotExpr); ok {
		return true
	}
	return false
}

func shuntOperator(expr ast.Expr, operatorStack StackExpr, exprStack StackExpr) (StackExpr, StackExpr, bool) {
	var y, x ast.Expr
	if not, ok := expr.(*ast.NotExpr); ok {
		if len(exprStack) < 1 {
			return operatorStack, exprStack, false
		}
		operatorStack, expr = operatorStack.pop()
		exprStack, x = exprStack.pop()
		not.SetX(x)
	} else if operator, ok := expr.(ast.OperatorExpr); ok {
		if len(exprStack) < 2 {
			return operatorStack, exprStack, false
		}
		operatorStack, expr = operatorStack.pop()
		exprStack, y = exprStack.pop()
		exprStack, x = exprStack.pop()
		operator.SetY(y)
		operator.SetX(x)
	}
	exprStack = exprStack.Push(expr)
	return operatorStack, exprStack, true
}
