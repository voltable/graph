package interpreter

import (
	"fmt"
	"github.com/voltable/graph/expressions/stack"
	"reflect"
)

type AddInstruction struct {
	add Operation
}

func NewAddInstruction(i interface{}) (*AddInstruction, error) {
	var op Operation

	var kind reflect.Kind

	if k, ok := i.(reflect.Kind); ok {
		kind = k
	} else {
		kind = reflect.TypeOf(i).Kind()
	}

	switch kind {
	case reflect.Int8:
		op = addInt8
	case reflect.Uint8:
		op = addUInt8
	case reflect.Int16:
		op = addInt16
	case reflect.Uint16:
		op = addUInt16
	case reflect.Int32:
		op = addInt32
	case reflect.Uint32:
		op = addUInt32
	case reflect.Int64:
		op = addInt64
	case reflect.Uint64:
		op = addUInt64
	case reflect.Int:
		op = addInt
	case reflect.Uint:
		op = addUInt
	case reflect.Float32:
		op = addFloat32
	case reflect.Float64:
		op = addFloat64
	case reflect.Complex64:
		op = addComplex64
	case reflect.Complex128:
		op = addComplex128
	default:
		return nil, fmt.Errorf("unsupport type %T", i)
	}

	return &AddInstruction{
		add: op,
	}, nil
}

func (s *AddInstruction) Run(stack *stack.Stack)  {
	Fetch(stack, s.add)
}

func addInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int8); ok {
		if r, ok := right.(int8); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addUInt8(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint8); ok {
		if r, ok := right.(uint8); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int16); ok {
		if r, ok := right.(int16); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addUInt16(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint16); ok {
		if r, ok := right.(uint16); ok {
			stack.Data[index-2] = l + r
		}
	}
}
func addInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int32); ok {
		if r, ok := right.(int32); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addUInt32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint32); ok {
		if r, ok := right.(uint32); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int64); ok {
		if r, ok := right.(int64); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addUInt64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint64); ok {
		if r, ok := right.(uint64); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addUInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(uint); ok {
		if r, ok := right.(uint); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addFloat32(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(float32); ok {
		if r, ok := right.(float32); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addFloat64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addComplex64(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(complex64); ok {
		if r, ok := right.(complex64); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func addComplex128(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(complex128); ok {
		if r, ok := right.(complex128); ok {
			stack.Data[index-2] = l + r
		}
	}
}
