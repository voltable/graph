package ast

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// MapProjectionStmt begins with the variable bound to the graph entity to be projected from, and contains a body of comma-separated map elements, enclosed by { and }.
type MapProjectionStmt struct {
	Variable string
	Elements []MapElementStmt
}

// NewMapProjectionStmt
func NewMapProjectionStmt(v string, e ...MapElementStmt) *MapProjectionStmt {
	return &MapProjectionStmt{Variable: v, Elements: e}
}

func (m *MapProjectionStmt) Interpret(variable string, prop vertices.Properties) []interface{} {
	arr := make([]interface{}, 0)
	if variable == m.Variable || m.Variable == "*" {
		for _, e := range m.Elements {
			arr = append(arr, e.Interpret(variable, prop))
		}
	}
	return arr
}

// MapElementStmt projects one or more key-value pairs to the map projection.
type MapElementStmt interface {
	mapElement()
	Interpret(string, vertices.Properties) interface{}
}

// MapProperty selector - Projects the property name as the key, and the value from the map_variable as the value for the projection.
// .name AS Test
type MapProperty struct {
	Key   string
	Alias string
}

func (*MapProperty) mapElement() {}

func (m *MapProperty) Interpret(variable string, prop vertices.Properties) interface{} {
	key := m.Key
	if m.Alias != StringEmpty {
		key = m.Alias
	}

	return struct {
		key   string
		value interface{}
	}{
		key:   key,
		value: prop.Property(m.Key),
	}
}

// MapLiteral This is a key-value pair, with the value being arbitrary expression
// key: <expression>
type MapLiteral struct {
	Key        string
	Expression Expr
	Alias      string
}

func (*MapLiteral) mapElement() {}

func (m *MapLiteral) Interpret(variable string, prop vertices.Properties) interface{} {

	if inter, ok := m.Expression.(InterpretExpr); ok {

		key := m.Key
		if m.Alias != StringEmpty {
			key = m.Alias
		}

		return struct {
			key   string
			value interface{}
		}{
			key:   key,
			value: inter.Interpret(variable, prop),
		}
	}

	return nil
}

// MapVariable Projects a variable, with the variable name as the key, and the value the variable is pointing to as the value of the projection. Its syntax is just the variable.
type MapVariable struct {
	Key   string
	Alias string
}

func (*MapVariable) mapElement() {}

func (m *MapVariable) Interpret(variable string, prop vertices.Properties) interface{} {
	// todo
	return nil
}

// MapAll All-properties selector - projects all key-value pairs from the map_variable value.
type MapAll struct {
}

func (*MapAll) mapElement() {}

func (m *MapAll) Interpret(variable string, prop vertices.Properties) interface{} {
	return prop
}
