package expressions_test

import (
	"github.com/voltable/graph/expressions"
	"testing"
)

type arithmeticExpression func(left expressions.TerminalExpression, right expressions.TerminalExpression) *expressions.BinaryArithmeticExpression

func TestBinaryArithmeticExpression_String(t *testing.T) {
	type args struct {
		left            expressions.TerminalExpression
		right           expressions.TerminalExpression
		expression		arithmeticExpression
	}
	tests := []struct {
		name   string
		args args
		want   string
	}{
		{
			name: "Add",
			args:args{
				left: expressions.Constant(1),
				right:  expressions.Constant(1),
				expression : expressions.Add,
			},
			want: "(1 + 1)",
		},
		{
			name: "Divide",
			args:args{
				left: expressions.Constant(1),
				right:  expressions.Constant(1),
				expression : expressions.Divide,
			},
			want: "(1 / 1)",
		},
		{
			name: "Modulo",
			args:args{
				left: expressions.Constant(1),
				right:  expressions.Constant(1),
				expression : expressions.Modulo,
			},
			want: "(1 % 1)",
		},
		{
			name: "Multiply",
			args:args{
				left:  expressions.Constant(1) ,
				right: expressions.Constant(1) ,
				expression : expressions.Multiply,
			},
			want: "(1 * 1)",
		},
		{
			name: "Power",
			args:args{
				left:  expressions.Constant(1),
				right: expressions.Constant(1) ,
				expression : expressions.Power,
			},
			want: "(1 ^ 1)",
		},
		{
			name: "Subtract",
			args:args{
				left:  expressions.Constant(1),
				right:  expressions.Constant(1),
				expression : expressions.Subtract,
			},
			want: "(1 - 1)",
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.args.expression(tt.args.left, tt.args.right)
			
			if got := got.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}