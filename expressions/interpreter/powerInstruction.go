package interpreter

import (
	"fmt"

	"github.com/voltable/graph/expressions/stack"
)

type PowerInstruction struct {
	power Operation
}

func NewPowerInstruction(i interface{}) (*PowerInstruction, error) {
	var op Operation

	switch i.(type) {
	case int8:
		op = powerInt8
	case uint8:
		op = powerUInt8
	case int16:
		op = powerInt16
	case uint16:
		op = powerUInt16
	case int32:
		op = powerInt32
	case uint32:
		op = powerUInt32
	case int64:
		op = powerInt64
	case uint64:
		op = powerUInt64
	case int:
		op = powerInt
	case uint:
		op = powerUInt
	default:
		return nil, fmt.Errorf("unsupport type %T", i)
	}

	return &PowerInstruction{
		power: op,
	}, nil
}

func (s *PowerInstruction) Run(stack *stack.Stack) {
	Fetch(stack, s.power)
}

func powerInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int8); ok {
		if r, ok := right.(int8); ok {
			stack.Data[index-2] = l ^ r
		}
	}
}

func powerUInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint8); ok {
		if r, ok := right.(uint8); ok {
			stack.Data[index-2] = l ^ r
		}
	}
}

func powerInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int16); ok {
		if r, ok := right.(int16); ok {
			stack.Data[index-2] = l ^ r
		}
	}
}

func powerUInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint16); ok {
		if r, ok := right.(uint16); ok {
			stack.Data[index-2] = l ^ r
		}
	}
}
func powerInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int32); ok {
		if r, ok := right.(int32); ok {
			stack.Data[index-2] = l ^ r
		}
	}
}

func powerUInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint32); ok {
		if r, ok := right.(uint32); ok {
			stack.Data[index-2] = l ^ r
		}
	}
}

func powerInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int64); ok {
		if r, ok := right.(int64); ok {
			stack.Data[index-2] = l ^ r
		}
	}
}

func powerUInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint64); ok {
		if r, ok := right.(uint64); ok {
			stack.Data[index-2] = l ^ r
		}
	}
}

func powerInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			stack.Data[index-2] = l ^ r
		}
	}
}

func powerUInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint); ok {
		if r, ok := right.(uint); ok {
			stack.Data[index-2] = l ^ r
		}
	}
}
