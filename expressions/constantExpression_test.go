package expressions_test

import (
"github.com/voltable/graph/expressions"
"testing"
)

func TestConstantExpression_String(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string

	}{
		{
			name: "Value",
			args:args{i:3.5},
			want: "3.5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expressions.Constant(tt.args.i)
			if got.String() != tt.want {
				t.Errorf("Constant() = %v, want %v", got, tt.want)
			}
		})
	}
}