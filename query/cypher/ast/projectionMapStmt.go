package ast

import (
	"github.com/RossMerr/Caudex.Graph/keyvalue"
)

// ProjectionMapStmt begins with the variable bound to the graph entity to be projected from, and contains a body of comma-separated map elements, enclosed by { and }.
type ProjectionMapStmt struct {
	Variable string
	Elements []ProjectionMapElementStmt
}

// NewProjectionMapStmt
func NewProjectionMapStmt(v string, e ...ProjectionMapElementStmt) *ProjectionMapStmt {
	return &ProjectionMapStmt{Variable: v, Elements: e}
}

func (m *ProjectionMapStmt) Interpret(variable string, prop *keyvalue.KeyValue) []interface{} {
	arr := make([]interface{}, 0)
	if variable == m.Variable || m.Variable == "*" {
		for _, e := range m.Elements {
			arr = append(arr, e.Interpret(variable, prop))
		}
	}
	return arr
}

// ProjectionMapElementStmt projects one or more key-value pairs to the map projection.
type ProjectionMapElementStmt interface {
	mapElement()
	Interpret(string, *keyvalue.KeyValue) interface{}
}

// ProjectionMapProperty selector - Projects the property name as the key, and the value from the map_variable as the value for the projection.
// .name AS Test
type ProjectionMapProperty struct {
	Key   string
	Alias string
}

func (*ProjectionMapProperty) mapElement() {}

func (m *ProjectionMapProperty) Interpret(variable string, prop *keyvalue.KeyValue) interface{} {
	key := m.Key
	if m.Alias != StringEmpty {
		key = m.Alias
	}

	return KeyValueStmt{
		Key:   key,
		Value: prop.Value.Unmarshal(),
	}
}

// ProjectionMapLiteral This is a key-value pair, with the value being arbitrary expression
// key: <expression>
type ProjectionMapLiteral struct {
	Key        string
	Expression Expr
	Alias      string
}

func (*ProjectionMapLiteral) mapElement() {}

func (m *ProjectionMapLiteral) Interpret(variable string, prop *keyvalue.KeyValue) interface{} {

	key := m.Key
	if m.Alias != StringEmpty {
		key = m.Alias
	}
	if inter, ok := m.Expression.(InterpretExpr); ok {
		return KeyValueStmt{
			Key:   key,
			Value: inter.Interpret(variable, prop),
		}
	}

	return KeyValueStmt{
		Key:   key,
		Value: false,
	}
}

// ProjectionMapVariable Projects a variable, with the variable name as the key, and the value the variable is pointing to as the value of the projection. Its syntax is just the variable.
type ProjectionMapVariable struct {
	Key   string
	Alias string
}

func (*ProjectionMapVariable) mapElement() {}

func (m *ProjectionMapVariable) Interpret(variable string, prop *keyvalue.KeyValue) interface{} {
	// todo
	return nil
}

// ProjectionMapAll All-properties selector - projects all key-value pairs from the map_variable value.
type ProjectionMapAll struct {
}

func (*ProjectionMapAll) mapElement() {}

func (m *ProjectionMapAll) Interpret(variable string, prop *keyvalue.KeyValue) interface{} {
	return prop
}
