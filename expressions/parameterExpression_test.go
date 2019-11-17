package expressions_test

import (
	"github.com/voltable/graph/expressions"
	"reflect"
	"testing"
)

func TestParameterExpression_String(t *testing.T) {
	type args struct {
		i reflect.Kind
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
			args:args{i: reflect.Int, name: ""},
		},
		{
			name: "Named",
			want: "test",
			args:args{i: reflect.Int, name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := expressions.Parameter(tt.args.i, tt.args.name)
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}