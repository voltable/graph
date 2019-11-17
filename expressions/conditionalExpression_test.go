package expressions_test

import (
	"github.com/voltable/graph/expressions"
	"testing"
)

type conditionalExpression func(test, ifTrue, ifFalse expressions.Expression)  *expressions.ConditionalExpression


func TestConditionalExpression_String(t *testing.T) {
	type args struct {
		test            expressions.TerminalExpression
		ifTrue           expressions.TerminalExpression
		ifFalse           expressions.TerminalExpression
		expression		conditionalExpression
	}
	tests := []struct {
		name   string
		args args
		want   string
	}{
		{
			name: "Condition",
			args:args{
				test: expressions.Constant(true) ,
				ifTrue:  expressions.Constant(1),
				ifFalse:  expressions.Constant(1) ,
				expression : expressions.Condition,
			},
			want: "IF(true, 1, 1)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.args.expression(tt.args.test, tt.args.ifTrue, tt.args.ifFalse)

			if got := got.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}