package ast

import "github.com/RossMerr/Caudex.Graph"

var _ TerminalExpr = (*PropertyStmt)(nil)
var _ InterpretExpr = (*PropertyStmt)(nil)

// PropertyStmt represents a node property.
type PropertyStmt struct {
	Variable string
	Value    string
}

func (PropertyStmt) interpretNode() {}
func (PropertyStmt) exprNode()      {}

func (p *PropertyStmt) GetValue() interface{} {
	return p.Value
}

func (p *PropertyStmt) SetValue(x interface{}) {
	if s, ok := x.(string); ok {
		p.Value = s
	}
}

func (p *PropertyStmt) Interpret(variable string, prop graph.Properties) interface{} {
	if p.Variable == variable {
		return prop.Property(p.Value)
	}
	return false
}
