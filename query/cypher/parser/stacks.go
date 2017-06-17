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

// Pop removes the last item on the StackExpr and returns it
func (s StackExpr) Pop() (StackExpr, ast.Expr, bool) {
	l := len(s)
	if l > 0 {
		return s[:l-1], s[l-1], true
	}
	return s, nil, false
}

// Top returns the last item on the StackExpr without removing it
func (s StackExpr) Top() (ast.Expr, bool) {
	l := len(s)
	if l > 0 {
		return s[l-1], true
	}
	return nil, false
}

// Shunt builds up the AST by Shunting the stack
func (s StackExpr) Shunt() (ast.Expr, error) {
	var item ast.Expr

	exprStack := make(StackExpr, 0)
	operatorStack := make(StackExpr, 0)
	notStack := make(StackExpr, 0)

	for len(s) > 0 {
		s, item, _ = s.Pop()
		if _, ok := item.(*ast.ParenthesesExpr); ok {
			// if p.Parentheses == ast.RPAREN {
			// 	operatorStack = operatorStack.Push(item)
			// } else { // LPAREN
			// 	var x ast.Expr
			// 	var y ast.Expr
			// 	var expr ast.Expr
			// 	operatorStack, expr, _ = operatorStack.Pop()
			// 	for expr != nil {
			// 		if p, ok := expr.(*ast.ParenthesesExpr); ok && p.Parentheses == ast.RPAREN {
			// 			break
			// 		} else {
			// 			exprStack, x, _ = exprStack.Pop()
			// 			exprStack, y, _ = exprStack.Pop()
			// 			if operator, ok := expr.(ast.OperatorExpr); ok {
			// 				operator.SetX(x)
			// 				operator.SetY(y)
			// 				exprStack = exprStack.Push(expr)
			// 			}
			// 		}
			// 		operatorStack, expr, _ = operatorStack.Pop()

			// 	}
			// }
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
		var x ast.Expr
		var y ast.Expr
		operatorStack, expr, _ = operatorStack.Pop()
		if operator, ok := expr.(ast.OperatorExpr); ok {
			exprStack, x, _ = exprStack.Pop()
			operator.SetX(x)
			exprStack, y, _ = exprStack.Pop()
			operator.SetY(y)
			fmt.Printf("%s went on exprStack \n", expr)
			exprStack = exprStack.Push(expr)

		}
	}

	var result ast.Expr

	exprStack, result, _ = exprStack.Pop()
	return result, nil
}

func shuntOperator(item ast.Expr, operatorStack StackExpr, exprStack StackExpr, notStack StackExpr) (StackExpr, StackExpr, StackExpr) {
	var x ast.Expr
	var y ast.Expr
	//fmt.Printf("Precedence first: %s (%s), second: %s (%s) \n", strconv.Itoa(ast.Precedence(expr)), expr, strconv.Itoa(ast.Precedence(item)), item)

	for expr, _ := operatorStack.Top(); expr != nil && ast.Precedence(expr) <= ast.Precedence(item); expr, _ = operatorStack.Top() {
		//	fmt.Printf(" first: %s (%s), second: %s (%s) \n", strconv.Itoa(ast.Precedence(expr)), expr, strconv.Itoa(ast.Precedence(item)), item)

		operatorStack, expr, _ = operatorStack.Pop()
		if operator, ok := expr.(ast.OperatorExpr); ok {

			exprStack, x, _ = exprStack.Pop()
			fmt.Printf("pop 1 %s \n", x)
			operator.SetX(x)

			exprStack, y, _ = exprStack.Pop()
			fmt.Printf("pop 2 %s \n", y)

			operator.SetY(y)
			fmt.Printf("%s went on exprStack \n", expr)

			// If we find anything on the notStack we should make the operator a child of it
			if len(notStack) > 0 {
				var n ast.Expr
				notStack, n, _ = notStack.Pop()
				if not, ok := n.(ast.OperatorExpr); ok {
					not.SetX(expr)
					exprStack = exprStack.Push(n)
				}
			} else {
				exprStack = exprStack.Push(expr)
			}
		}
	}

	fmt.Printf("%s went on operatorStack \n", item)

	operatorStack = operatorStack.Push(item)

	return operatorStack, exprStack, notStack
}
