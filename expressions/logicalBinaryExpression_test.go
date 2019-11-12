package expressions_test

import (
	"github.com/voltable/graph/expressions"
	"testing"
)
type logicalExpression func(left expressions.TerminalExpression, right expressions.TerminalExpression) (*expressions.LogicalBinaryExpression, error)


func TestLogicalBinaryExpression_String(t *testing.T) {
	type args struct {
		left            expressions.TerminalExpression
		right           expressions.TerminalExpression
		expression		logicalExpression
	}
	tests := []struct {
		name   string
		args args
		want   string
	}{
		{
			name: "Equal",
			args:args{
				left: func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				right:  func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				expression : expressions.Equal,
			},
			want: "(1 = 1)",
		},
		{
			name: "NotEqual",
			args:args{
				left: func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				right:  func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				expression : expressions.NotEqual,
			},
			want: "(1 <> 1)",
		},
		{
			name: "LessThan",
			args:args{
				left: func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				right:  func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				expression : expressions.LessThan,
			},
			want: "(1 < 1)",
		},
		{
			name: "LessThanOrEqual",
			args:args{
				left: func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				right:  func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				expression : expressions.LessThanOrEqual,
			},
			want: "(1 <= 1)",
		},
		{
			name: "GreaterThan",
			args:args{
				left: func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				right:  func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				expression : expressions.GreaterThan,
			},
			want: "(1 > 1)",
		},
		{
			name: "GreaterThanOrEqual",
			args:args{
				left: func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				right:  func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				expression : expressions.GreaterThanOrEqual,
			},
			want: "(1 >= 1)",
		},
		{
			name: "IsNil",
			args:args{
				left: func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				right:  func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				expression : expressions.IsNil,
			},
			want: "(1 IS NULL 1)",
		},
		{
			name: "IsNotNil",
			args:args{
				left: func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				right:  func() expressions.TerminalExpression {
					c, _ := expressions.Constant(1)
					return c
				}() ,
				expression : expressions.IsNotNil,
			},
			want: "(1 IS NOT NULL 1)",
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