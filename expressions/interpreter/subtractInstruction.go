package interpreter

import (
	"fmt"

	"github.com/voltable/graph/expressions/stack"
)

type SubtractInstruction struct {
	subtract operation
}

func NewSubtractInstruction(i interface{}) (*SubtractInstruction, error) {
	var op operation

	switch i.(type) {
	case int8:
		op = subtractInt8
	case uint8:
		op = subtractUInt8
	case int16:
		op = subtractInt16
	case uint16:
		op = subtractUInt16
	case int32:
		op = subtractInt32
	case uint32:
		op = subtractUInt32
	case int64:
		op = subtractInt64
	case uint64:
		op = subtractUInt64
	case int:
		op = subtractInt
	case uint:
		op = subtractUInt
	case float32:
		op = subtractFloat32
	case float64:
		op = subtractFloat64
	case complex64:
		op = subtractComplex64
	case complex128:
		op = subtractComplex128
	default:
		return nil, fmt.Errorf("unsupport type %T", i)
	}

	return &SubtractInstruction{
		subtract: op,
	}, nil
}

func (s *SubtractInstruction) Run(stack *stack.Stack) error {
	return fetch(stack, s.subtract)
}

func subtractInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int8); ok {
		if r, ok := right.(int8); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractUInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint8); ok {
		if r, ok := right.(uint8); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int16); ok {
		if r, ok := right.(int16); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractUInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint16); ok {
		if r, ok := right.(uint16); ok {
			stack.Data[index-2] = l - r
		}
	}
}
func subtractInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int32); ok {
		if r, ok := right.(int32); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractUInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint32); ok {
		if r, ok := right.(uint32); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int64); ok {
		if r, ok := right.(int64); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractUInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint64); ok {
		if r, ok := right.(uint64); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractUInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint); ok {
		if r, ok := right.(uint); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractFloat32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(float32); ok {
		if r, ok := right.(float32); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractFloat64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractComplex64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(complex64); ok {
		if r, ok := right.(complex64); ok {
			stack.Data[index-2] = l - r
		}
	}
}

func subtractComplex128(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(complex128); ok {
		if r, ok := right.(complex128); ok {
			stack.Data[index-2] = l - r
		}
	}
}
