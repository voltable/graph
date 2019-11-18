package interpreter

import "github.com/voltable/graph/expressions/stack"

type operation func(left interface{}, right interface{}, stack *stack.Stack, index int)


func fetch(stack *stack.Stack, op operation) error {
	index := stack.Index
	left := stack.Data[index-2]
	if left != nil {
		right := stack.Data[index-1]
		if right == nil {
			stack.Data[index-1] = nil
		} else {
			op(left, right, stack, index)
		}
	}

	stack.Index = index - 1

	return nil
}