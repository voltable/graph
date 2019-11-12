package expressions_test

import (
	"github.com/voltable/graph/expressions"
	"testing"
)

type booleanExpression func(left expressions.TerminalExpression, right expressions.TerminalExpression) (*expressions.BooleanExpression, error)


func TestBooleanExpression_String(t *testing.T) {
	type args struct {
		left            expressions.TerminalExpression
		right           expressions.TerminalExpression
		expression		booleanExpression
	}
	tests := []struct {
		name   string
		args args
		want   string
	}{
		{
			name: "And",
			args:args{
				left: func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				right:  func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				expression : expressions.And,
			},
			want: "(1 & 1)",
		},
		{
			name: "Or",
			args:args{
				left: func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				right:  func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				expression : expressions.Or,
			},
			want: "(1 | 1)",
		},
		{
			name: "Xor",
			args:args{
				left: func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				right:  func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				expression : expressions.Xor,
			},
			want: "(1 ^ 1)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := tt.args.expression(tt.args.left, tt.args.right)
			if got := got.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}