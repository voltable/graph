package ast

// MapProjectionStmt begins with the variable bound to the graph entity to be projected from, and contains a body of comma-separated map elements, enclosed by { and }.
type MapProjectionStmt struct {
	Variable string
	Elements []MapElementStmt
}

// NewMapProjectionStmt
func NewMapProjectionStmt(v string, e ...MapElementStmt) *MapProjectionStmt {
	return &MapProjectionStmt{Variable: v, Elements: e}
}

// MapElementStmt projects one or more key-value pairs to the map projection.
type MapElementStmt interface {
	mapElement()
}

// MapProperty selector - Projects the property name as the key, and the value from the map_variable as the value for the projection.
type MapProperty struct {
	Key string
}

func (*MapProperty) mapElement() {}

// MapLiteral This is a key-value pair, with the value being arbitrary expression
type MapLiteral struct {
	Key        string
	Expression Expr
}

func (*MapLiteral) mapElement() {}

// MapVariable Projects a variable, with the variable name as the key, and the value the variable is pointing to as the value of the projection. Its syntax is just the variable.
type MapVariable struct {
	Key string
}

func (*MapVariable) mapElement() {}

// MapAll All-properties selector - projects all key-value pairs from the map_variable value.
type MapAll struct {
}

func (*MapAll) mapElement() {}
