package expressions_test

import (
	"github.com/voltable/graph/expressions"
	"reflect"
	"testing"
)

func TestLambdaExpression_String(t *testing.T) {
	paramExpr := expressions.Parameter(reflect.Int, "arg")

	type args struct {
		expression  expressions.Expression
		parameters []*expressions.ParameterExpression
	}
	tests := []struct {
		name   string
		args args
		want   string
	}{
		{
			name: "Lambda",
			args: args{
				expressions.Add(paramExpr, expressions.Constant(1)),
				func() []*expressions.ParameterExpression {
					parameters := make([]*expressions.ParameterExpression, 0)
					parameters = append(parameters, paramExpr)
					return parameters
				}(),
			},
			want: "arg => (arg + 1)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expressions.Lambda(tt.args.expression, tt.args.parameters...)
			if got := got.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}