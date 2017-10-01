package ast

import "github.com/RossMerr/Caudex.Graph/vertices"

var _ TerminalExpr = (*Ident)(nil)

// Ident used to hold anu object or nil
type Ident struct {
	Data interface{} // denoted object; or nil
}

func (Ident) exprNode()      {}
func (Ident) interpretNode() {}

func (p *Ident) GetValue() interface{} {
	return p.Data
}

func (p *Ident) SetValue(x interface{}) {
	p.Data = x
}

func (p *Ident) Interpret(variable string, vertex *vertices.Vertex) interface{} {
	return p.Data
}
