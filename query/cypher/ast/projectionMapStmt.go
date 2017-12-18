package ast

import (
	"github.com/RossMerr/Caudex.Graph"
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

func (m *ProjectionMapStmt) Interpret(variable string, prop graph.Properties) []interface{} {
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
	Interpret(string, graph.Properties) interface{}
}

// ProjectionMapProperty selector - Projects the property name as the key, and the value from the map_variable as the value for the projection.
// .name AS Test
type ProjectionMapProperty struct {
	Key   string
	Alias string
}

func (*ProjectionMapProperty) mapElement() {}

func (m *ProjectionMapProperty) Interpret(variable string, prop graph.Properties) interface{} {
	key := m.Key
	if m.Alias != StringEmpty {
		key = m.Alias
	}

	return graph.KeyValue{
		Key:   key,
		Value: prop.Property(m.Key),
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

func (m *ProjectionMapLiteral) Interpret(variable string, prop graph.Properties) interface{} {

	key := m.Key
	if m.Alias != StringEmpty {
		key = m.Alias
	}
	if inter, ok := m.Expression.(InterpretExpr); ok {
		return graph.KeyValue{
			Key:   key,
			Value: inter.Interpret(variable, prop),
		}
	}

	return graph.KeyValue{
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

func (m *ProjectionMapVariable) Interpret(variable string, prop graph.Properties) interface{} {
	// todo
	return nil
}

// ProjectionMapAll All-properties selector - projects all key-value pairs from the map_variable value.
type ProjectionMapAll struct {
}

func (*ProjectionMapAll) mapElement() {}

func (m *ProjectionMapAll) Interpret(variable string, prop graph.Properties) interface{} {
	return prop
}
