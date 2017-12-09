package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_MapPropertyInterpret(t *testing.T) {
	var tests = []struct {
		expr     *ast.MapProperty
		variable string
		prop     graph.Properties
		result   interface{}
	}{
		{
			expr: &ast.MapProperty{Key: "name"},
			prop: func() graph.Properties {
				v, _ := graph.NewVertex()
				v.SetProperty("name", "john smith")
				return v
			}(),
			result: func() interface{} {
				kv := graph.KeyValue{
					Key:   "name",
					Value: "john smith",
				}
				return kv
			}(),
		},
		{
			expr: &ast.MapProperty{Key: "name", Alias: "alias"},
			prop: func() graph.Properties {
				v, _ := graph.NewVertex()
				v.SetProperty("name", "john smith")
				return v
			}(),
			result: func() interface{} {
				kv := graph.KeyValue{
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

func Test_MapLiteralInterpret(t *testing.T) {
	var tests = []struct {
		expr     *ast.MapLiteral
		variable string
		prop     graph.Properties
		result   interface{}
	}{
		{
			expr: &ast.MapLiteral{Key: "name", Expression: &ast.ComparisonExpr{Comparison: expressions.EQ}},
			prop: func() graph.Properties {
				v, _ := graph.NewVertex()
				v.SetProperty("name", "john smith")
				return v
			}(),
			result: func() interface{} {
				kv := graph.KeyValue{
					Key:   "name",
					Value: true,
				}
				return kv
			}(),
		},
		{
			expr: &ast.MapLiteral{Key: "name", Alias: "alias", Expression: &ast.ComparisonExpr{Comparison: expressions.EQ}},
			prop: func() graph.Properties {
				v, _ := graph.NewVertex()
				v.SetProperty("name", "john smith")
				return v
			}(),
			result: func() interface{} {
				kv := graph.KeyValue{
					Key:   "alias",
					Value: true,
				}
				return kv
			}(),
		},
		{
			expr: &ast.MapLiteral{Key: "name", Expression: nil},
			prop: func() graph.Properties {
				v, _ := graph.NewVertex()
				v.SetProperty("name", "john smith")
				return v
			}(),
			result: func() interface{} {
				kv := graph.KeyValue{
					Key:   "name",
					Value: false,
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

func Test_MapAllInterpret(t *testing.T) {

	v, _ := graph.NewVertex()
	v.SetProperty("name", "john smith")

	var tests = []struct {
		expr     *ast.MapAll
		variable string
		prop     graph.Properties
		result   interface{}
	}{
		{
			expr: &ast.MapAll{},
			prop: func() graph.Properties {
				return v
			}(),
			result: func() interface{} {
				return v

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
