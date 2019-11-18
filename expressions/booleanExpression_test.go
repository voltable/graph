package expressions_test

import (
	"github.com/voltable/graph/expressions"
	"testing"
)

type booleanExpression func(left expressions.TerminalExpression, right expressions.TerminalExpression) *expressions.BooleanExpression


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
				left: expressions.Constant(1) ,
				right: expressions.Constant(1),
				expression : expressions.And,
			},
			want: "(1 & 1)",
		},
		{
			name: "Or",
			args:args{
				left: expressions.Constant(1) ,
				right:  expressions.Constant(1),
				expression : expressions.Or,
			},
			want: "(1 | 1)",
		},
		{
			name: "Xor",
			args:args{
				left:  expressions.Constant(1),
				right: expressions.Constant(1) ,
				expression : expressions.Xor,
			},
			want: "(1 ^ 1)",
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




