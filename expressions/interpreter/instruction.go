package interpreter

import "github.com/voltable/graph/expressions/stack"

type Instruction interface {
	Run(stack *stack.Stack) error
}
