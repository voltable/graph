package expressions_test

import (
	"github.com/voltable/graph/expressions"
	"testing"
)

func TestParameterExpression_String(t *testing.T) {
	type args struct {
		i interface{}
		name string
	}
	tests := []struct {
		name   string
		args args
		want   string
	}{
		{
			name: "No name",
			want: "Param_0",
			args:args{i: "1", name: ""},
		},
		{
			name: "Named",
			want: "test",
			args:args{i: "1", name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, _ := expressions.Parameter(tt.args.i, tt.args.name)
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}