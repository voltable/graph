package interpreter_test

import (
	"github.com/voltable/graph/expressions/interpreter"
	"github.com/voltable/graph/expressions/stack"
	"testing"
)

func addInt(left interface{}, right interface{}, stack *stack.Stack, index int) {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			stack.Data[index-2] = l + r
		}
	}
}

func Test_fetch(t *testing.T) {
	type args struct {
		stack *stack.Stack
		op    interpreter.Operation
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{

		{
			name:"add",
			args:args{
				stack: func() *stack.Stack{
					s := &stack.Stack{}
					s.Push(1)
					s.Push(2)
					return s
				}(),
				op:  addInt,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interpreter.Fetch(tt.args.stack, tt.args.op)
			got := tt.args.stack.Pop()
			if tt.want !=  got{
				t.Errorf("Fetch() = %v, want %v", got, tt.want)
			}

		})
	}
}