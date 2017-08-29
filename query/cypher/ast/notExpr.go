package ast

import "github.com/RossMerr/Caudex.Graph/vertices"

// NotExpr a bitwise complement or logical negation operation (Not a).
type NotExpr struct {
	value Expr // left operand
}

var _ TerminalExpr = (*NotExpr)(nil)

func (NotExpr) exprNode() {}

func (b *NotExpr) SetValue(left Expr) {
	b.value = left
}

func (b *NotExpr) GetValue() Expr {
	return b.value
}

// Interpret runs the NotExpr over a Vertex and VertexPatn to check for a match
func (b *NotExpr) Interpret(vertex *vertices.Vertex, pattern *VertexPatn) bool {
	return false
}

// NotPrecedence returns the precedence (order of importance)
func NotPrecedence(item NotExpr) int {
	return 13
}
