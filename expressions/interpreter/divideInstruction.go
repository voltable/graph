package interpreter

import (
	"fmt"

	"github.com/voltable/graph/expressions/stack"
)

type DivideInstruction struct {
	divide Operation
}

func NewDivideInstruction(i interface{}) (*DivideInstruction, error) {
	var op Operation

	switch i.(type) {
	case int8:
		op = divideInt8
	case uint8:
		op = divideUInt8
	case int16:
		op = divideInt16
	case uint16:
		op = divideUInt16
	case int32:
		op = divideInt32
	case uint32:
		op = divideUInt32
	case int64:
		op = divideInt64
	case uint64:
		op = divideUInt64
	case int:
		op = divideInt
	case uint:
		op = divideUInt
	case float32:
		op = divideFloat32
	case float64:
		op = divideFloat64
	case complex64:
		op = divideComplex64
	case complex128:
		op = divideComplex128
	default:
		return nil, fmt.Errorf("unsupport type %T", i)
	}

	return &DivideInstruction{
		divide: op,
	}, nil
}

func (s *DivideInstruction) Run(stack *stack.Stack)  {
	Fetch(stack, s.divide)
}

func divideInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int8); ok {
		if r, ok := right.(int8); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideUInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint8); ok {
		if r, ok := right.(uint8); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int16); ok {
		if r, ok := right.(int16); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideUInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint16); ok {
		if r, ok := right.(uint16); ok {
			stack.Data[index-2] = l / r
		}
	}
}
func divideInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int32); ok {
		if r, ok := right.(int32); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideUInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint32); ok {
		if r, ok := right.(uint32); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int64); ok {
		if r, ok := right.(int64); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideUInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint64); ok {
		if r, ok := right.(uint64); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideUInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint); ok {
		if r, ok := right.(uint); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideFloat32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(float32); ok {
		if r, ok := right.(float32); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideFloat64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideComplex64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(complex64); ok {
		if r, ok := right.(complex64); ok {
			stack.Data[index-2] = l / r
		}
	}
}

func divideComplex128(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(complex128); ok {
		if r, ok := right.(complex128); ok {
			stack.Data[index-2] = l / r
		}
	}
}
