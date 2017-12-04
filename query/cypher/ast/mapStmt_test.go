package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_MapPropertyInterpret(t *testing.T) {
	var tests = []struct {
		expr     *ast.MapProperty
		variable string
		prop     vertices.Properties
		result   interface{}
	}{
		{
			expr: &ast.MapProperty{Key: "name"},
			prop: func() vertices.Properties {
				v, _ := vertices.NewVertex()
				v.SetProperty("name", "john smith")
				return v
			}(),
			result: func() interface{} {
				kv := vertices.KeyValue{
					Key:   "name",
					Value: "john smith",
				}
				return kv
			}(),
		},
		{
			expr: &ast.MapProperty{Key: "name", Alias: "alias"},
			prop: func() vertices.Properties {
				v, _ := vertices.NewVertex()
				v.SetProperty("name", "john smith")
				return v
			}(),
			result: func() interface{} {
				kv := vertices.KeyValue{
					Key:   "alias",
					Value: "john smith",
				}
				return kv
			}(),
		},
	}

	for i, tt := range tests {
		result := tt.expr.Interpret(tt.variable, tt.prop)
		if result != tt.result {
			t.Errorf("%d.  %q: comparison mismatch:\n  exp=%v\n  got=%v\n\n", i, tt.expr, tt.result, result)
		}
	}

}
