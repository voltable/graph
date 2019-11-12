package expressions

import "github.com/voltable/graph/expressions/stack"

// InterpretExpression is the base interface for the NonTerminalExpression and TerminalExpression
type InterpretExpression interface {
	Expression
	Interpret(stack *stack.Stack) error
}


