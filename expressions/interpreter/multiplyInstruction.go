package interpreter

import (
	"fmt"

	"github.com/voltable/graph/expressions/stack"
)

type MultiplyInstruction struct {
	multiply operation
}

func NewMultiplyInstruction(i interface{}) (*MultiplyInstruction, error) {
	var op operation

	switch i.(type) {
	case int8:
		op = multiplyInt8
	case uint8:
		op = multiplyUInt8
	case int16:
		op = multiplyInt16
	case uint16:
		op = multiplyUInt16
	case int32:
		op = multiplyInt32
	case uint32:
		op = multiplyUInt32
	case int64:
		op = multiplyInt64
	case uint64:
		op = multiplyUInt64
	case int:
		op = multiplyInt
	case uint:
		op = multiplyUInt
	case float32:
		op = multiplyFloat32
	case float64:
		op = multiplyFloat64
	case complex64:
		op = multiplyComplex64
	case complex128:
		op = multiplyComplex128
	default:
		return nil, fmt.Errorf("unsupport type %T", i)
	}

	return &MultiplyInstruction{
		multiply: op,
	}, nil
}

func (s *MultiplyInstruction) Run(stack *stack.Stack) error {
	return fetch(stack, s.multiply)
}

func multiplyInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int8); ok {
		if r, ok := right.(int8); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyUInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint8); ok {
		if r, ok := right.(uint8); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int16); ok {
		if r, ok := right.(int16); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyUInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint16); ok {
		if r, ok := right.(uint16); ok {
			stack.Data[index-2] = l * r
		}
	}
}
func multiplyInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int32); ok {
		if r, ok := right.(int32); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyUInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint32); ok {
		if r, ok := right.(uint32); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int64); ok {
		if r, ok := right.(int64); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyUInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint64); ok {
		if r, ok := right.(uint64); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyUInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint); ok {
		if r, ok := right.(uint); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyFloat32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(float32); ok {
		if r, ok := right.(float32); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyFloat64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyComplex64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(complex64); ok {
		if r, ok := right.(complex64); ok {
			stack.Data[index-2] = l * r
		}
	}
}

func multiplyComplex128(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(complex128); ok {
		if r, ok := right.(complex128); ok {
			stack.Data[index-2] = l * r
		}
	}
}
