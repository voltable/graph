package expressions_test

import (
	"github.com/voltable/graph/expressions"
	"reflect"
	"testing"
)

func TestInvocation_String(t *testing.T) {
	num1 := expressions.Parameter(reflect.Int, "num1")
	num2 := expressions.Parameter(reflect.Int, "num2")

	type args struct {
		expression expressions.Expression
		parameters []*expressions.ParameterExpression
		left       expressions.TerminalExpression
		right      expressions.TerminalExpression
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Invoke",
			args: args{
				expression: expressions.GreaterThan(expressions.Add(num1, num2), expressions.Constant(1000)),
				parameters: func() []*expressions.ParameterExpression {
					parameters := make([]*expressions.ParameterExpression, 0)
					parameters = append(parameters, num1)
					parameters = append(parameters, num2)
					return parameters
				}(),
				left:  expressions.Constant(539),
				right: expressions.Constant(281),
			},
			want: "Invoke((num1, num2) => ((num1 + num2) > 1000), 539, 281)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expressions.Lambda(tt.args.expression, tt.args.parameters...)

			invoke := expressions.Invoke(got, tt.args.left, tt.args.right)

			if got := invoke.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
