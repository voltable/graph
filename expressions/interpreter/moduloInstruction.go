package interpreter

import (
	"fmt"

	"github.com/voltable/graph/expressions/stack"
)

type ModuloInstruction struct {
	modulo operation
}

func NewModuloInstruction(i interface{}) (*ModuloInstruction, error) {
	var op operation

	switch i.(type) {
	case int8:
		op = moduloInt8
	case uint8:
		op = moduloUInt8
	case int16:
		op = moduloInt16
	case uint16:
		op = moduloUInt16
	case int32:
		op = moduloInt32
	case uint32:
		op = moduloUInt32
	case int64:
		op = moduloInt64
	case uint64:
		op = moduloUInt64
	case int:
		op = moduloInt
	case uint:
		op = moduloUInt
	default:
		return nil, fmt.Errorf("unsupport type %T", i)
	}

	return &ModuloInstruction{
		modulo: op,
	}, nil
}

func (s *ModuloInstruction) Run(stack *stack.Stack) error {
	return fetch(stack, s.modulo)
}

func moduloInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int8); ok {
		if r, ok := right.(int8); ok {
			stack.Data[index-2] = l % r
		}
	}
}

func moduloUInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint8); ok {
		if r, ok := right.(uint8); ok {
			stack.Data[index-2] = l % r
		}
	}
}

func moduloInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int16); ok {
		if r, ok := right.(int16); ok {
			stack.Data[index-2] = l % r
		}
	}
}

func moduloUInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint16); ok {
		if r, ok := right.(uint16); ok {
			stack.Data[index-2] = l % r
		}
	}
}
func moduloInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int32); ok {
		if r, ok := right.(int32); ok {
			stack.Data[index-2] = l % r
		}
	}
}

func moduloUInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint32); ok {
		if r, ok := right.(uint32); ok {
			stack.Data[index-2] = l % r
		}
	}
}

func moduloInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int64); ok {
		if r, ok := right.(int64); ok {
			stack.Data[index-2] = l % r
		}
	}
}

func moduloUInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint64); ok {
		if r, ok := right.(uint64); ok {
			stack.Data[index-2] = l % r
		}
	}
}

func moduloInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			stack.Data[index-2] = l % r
		}
	}
}

func moduloUInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint); ok {
		if r, ok := right.(uint); ok {
			stack.Data[index-2] = l % r
		}
	}
}
