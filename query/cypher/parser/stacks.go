package parser

import (
	"fmt"

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

	for len(s) > 0 {
		s, item, _ = s.Pop()
		if p, ok := item.(*ast.ParenthesesExpr); ok {
			if p.Parentheses == ast.LPAREN {
				exprStack = operatorStack.Push(item)
			} else {
				var x ast.Expr
				var y ast.Expr
				expr, _ := operatorStack.Top()
				for expr != nil && expr.(*ast.ParenthesesExpr).Parentheses != ast.LPAREN {
					operatorStack, expr, _ = operatorStack.Pop()
					exprStack, x, _ = exprStack.Pop()
					exprStack, y, _ = exprStack.Pop()
					if operator, ok := expr.(ast.OperatorExpr); ok {
						operator.SetX(x)
						operator.SetY(y)
						exprStack = exprStack.Push(expr)

					}
				}
				// Pop the '(' off the operator stack.
				operatorStack, _, _ = operatorStack.Pop()
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
			var x ast.Expr
			var y ast.Expr
			//fmt.Printf("Precedence first: %s (%s), second: %s (%s) \n", strconv.Itoa(ast.Precedence(expr)), expr, strconv.Itoa(ast.Precedence(item)), item)

			for expr, _ := operatorStack.Top(); expr != nil && ast.Precedence(expr) <= ast.Precedence(item); expr, _ = operatorStack.Top() {
				//	fmt.Printf("first: %s (%s), second: %s (%s) \n", strconv.Itoa(ast.Precedence(expr)), expr, strconv.Itoa(ast.Precedence(item)), item)

				operatorStack, expr, _ = operatorStack.Pop()
				exprStack, x, _ = exprStack.Pop()
				fmt.Printf("pop 1 %s \n", x)
				exprStack, y, _ = exprStack.Pop()
				fmt.Printf("pop 2 %s \n", y)
				if operator, ok := expr.(ast.OperatorExpr); ok {
					operator.SetX(x)
					operator.SetY(y)
					fmt.Printf("%s went on exprStack \n", expr)
					exprStack = exprStack.Push(expr)

				}
			}

			fmt.Printf("%s went on operatorStack \n", item)
			operatorStack = operatorStack.Push(item)

		} else if _, ok := item.(*ast.BooleanExpr); ok {
			// Otherwise, the token is an operator (operator here includes both ComparisonExpr and BooleanExpr).
			var x ast.Expr
			var y ast.Expr
			//fmt.Printf("Precedence first: %s (%s), second: %s (%s) \n", strconv.Itoa(ast.Precedence(expr)), expr, strconv.Itoa(ast.Precedence(item)), item)

			for expr, _ := operatorStack.Top(); expr != nil && ast.Precedence(expr) <= ast.Precedence(item); expr, _ = operatorStack.Top() {
				//	fmt.Printf(" first: %s (%s), second: %s (%s) \n", strconv.Itoa(ast.Precedence(expr)), expr, strconv.Itoa(ast.Precedence(item)), item)

				operatorStack, expr, _ = operatorStack.Pop()
				exprStack, x, _ = exprStack.Pop()
				fmt.Printf("pop 1 %s \n", x)
				exprStack, y, _ = exprStack.Pop()
				fmt.Printf("pop 2 %s \n", y)
				if operator, ok := expr.(ast.OperatorExpr); ok {
					operator.SetX(x)
					operator.SetY(y)
					fmt.Printf("%s went on exprStack \n", expr)
					exprStack = exprStack.Push(expr)

				}
			}
			fmt.Printf("%s went on operatorStack \n", item)

			operatorStack = operatorStack.Push(item)

		}

	}

	for len(operatorStack) > 0 {
		var expr ast.Expr
		var x ast.Expr
		var y ast.Expr
		operatorStack, expr, _ = operatorStack.Pop()
		exprStack, x, _ = exprStack.Pop()
		exprStack, y, _ = exprStack.Pop()
		if operator, ok := expr.(ast.OperatorExpr); ok {
			operator.SetX(x)
			operator.SetY(y)
			fmt.Printf("%s went on exprStack \n", expr)
			exprStack = exprStack.Push(expr)

		}
	}

	var result ast.Expr

	exprStack, result, _ = exprStack.Pop()
	return result, nil
}
