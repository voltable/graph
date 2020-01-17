package expressions_test

import (
	"github.com/voltable/graph/expressions"
	"testing"
)

type logicalExpression func(left expressions.Expression, right expressions.Expression) *expressions.LogicalBinaryExpression

func TestLogicalBinaryExpression_String(t *testing.T) {
	type args struct {
		left       expressions.TerminalExpression
		right      expressions.TerminalExpression
		expression logicalExpression
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Equal",
			args: args{
				left:       expressions.Constant(1),
				right:      expressions.Constant(1),
				expression: expressions.Equal,
			},
			want: "(1 = 1)",
		},
		{
			name: "NotEqual",
			args: args{
				left:       expressions.Constant(1),
				right:      expressions.Constant(1),
				expression: expressions.NotEqual,
			},
			want: "(1 <> 1)",
		},
		{
			name: "LessThan",
			args: args{
				left:       expressions.Constant(1),
				right:      expressions.Constant(1),
				expression: expressions.LessThan,
			},
			want: "(1 < 1)",
		},
		{
			name: "LessThanOrEqual",
			args: args{
				left:       expressions.Constant(1),
				right:      expressions.Constant(1),
				expression: expressions.LessThanOrEqual,
			},
			want: "(1 <= 1)",
		},
		{
			name: "GreaterThan",
			args: args{
				left:       expressions.Constant(1),
				right:      expressions.Constant(1),
				expression: expressions.GreaterThan,
			},
			want: "(1 > 1)",
		},
		{
			name: "GreaterThanOrEqual",
			args: args{
				left:       expressions.Constant(1),
				right:      expressions.Constant(1),
				expression: expressions.GreaterThanOrEqual,
			},
			want: "(1 >= 1)",
		},
		{
			name: "IsNil",
			args: args{
				left:       expressions.Constant(1),
				right:      expressions.Constant(1),
				expression: expressions.IsNil,
			},
			want: "(1 IS NULL 1)",
		},
		{
			name: "IsNotNil",
			args: args{
				left:       expressions.Constant(1),
				right:      expressions.Constant(1),
				expression: expressions.IsNotNil,
			},
			want: "(1 IS NOT NULL 1)",
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
