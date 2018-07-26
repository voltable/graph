package ast_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/keyvalue"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

func Test_MapPropertyInterpret(t *testing.T) {
	var tests = []struct {
		expr     *ast.ProjectionMapProperty
		variable string
		prop     *keyvalue.KeyValue
		result   interface{}
	}{
		{
			expr: &ast.ProjectionMapProperty{Key: "name"},
			prop: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("John Smith", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
			}(),
			result: func() interface{} {
				kv := graph.KeyValue{
					Key:   "name",
					Value: "John Smith",
				}
				return kv
			}(),
		},
		{
			expr: &ast.ProjectionMapProperty{Key: "name", Alias: "alias"},
			prop: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("John Smith", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
			}(),
			result: func() interface{} {
				kv := graph.KeyValue{
					Key:   "alias",
					Value: "John Smith",
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
		expr     *ast.ProjectionMapLiteral
		variable string
		prop     *keyvalue.KeyValue
		result   interface{}
	}{
		{
			expr: &ast.ProjectionMapLiteral{Key: "name", Expression: &ast.ComparisonExpr{Comparison: expressions.EQ}},
			prop: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("John Smith", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
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
			expr: &ast.ProjectionMapLiteral{Key: "name", Alias: "alias", Expression: &ast.ComparisonExpr{Comparison: expressions.EQ}},
			prop: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("John Smith", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
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
			expr: &ast.ProjectionMapLiteral{Key: "name", Expression: nil},
			prop: func() *keyvalue.KeyValue {
				id, _ := graph.GenerateRandomUUID()
				x := keyvalue.NewKeyValue("John Smith", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))
				return x
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

	id, _ := graph.GenerateRandomUUID()
	v := keyvalue.NewKeyValue("John Smith", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte("name"))

	var tests = []struct {
		expr     *ast.ProjectionMapAll
		variable string
		prop     *keyvalue.KeyValue
		result   interface{}
	}{
		{
			expr: &ast.ProjectionMapAll{},
			prop: func() *keyvalue.KeyValue {
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
