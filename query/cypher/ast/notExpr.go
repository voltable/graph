package ast

import "github.com/RossMerr/Caudex.Graph/vertices"

// NotExpr a bitwise complement or logical negation operation (Not a).
type NotExpr struct {
	value InterpretExpr // left operand
}

var _ TerminalExpr = (*NotExpr)(nil)

func (NotExpr) exprNode()      {}
func (NotExpr) interpretNode() {}

func (b *NotExpr) SetValue(left InterpretExpr) {
	b.value = left
}

func (b *NotExpr) GetValue() InterpretExpr {
	return b.value
}

// Interpret runs the NotExpr over a Vertex to check for a match
func (b *NotExpr) Interpret(vertex *vertices.Vertex) bool {
	return false
}

// NotPrecedence returns the precedence (order of importance)
func NotPrecedence(item NotExpr) int {
	return 13
}
