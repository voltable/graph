package interpreter

import "github.com/voltable/graph/expressions/stack"

type Operation func(left interface{}, right interface{}, stack *stack.Stack, index int)


func Fetch(stack *stack.Stack, op Operation)  {
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
}