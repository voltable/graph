package ast

import (
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type BooleanExpr struct {
	expressions.Boolean
	Left  Expr // left operand
	Right Expr // right operand
}

// NotExpr a bitwise complement or logical negation operation (Not a).
type NotExpr struct {
	value Expr // left operand
}

func (BooleanExpr) exprNode() {}
func (NotExpr) exprNode()     {}

func (b *BooleanExpr) GetLeft() Expr {
	return b.Left
}

func (b *BooleanExpr) GetRight() Expr {
	return b.Right
}

func (b *BooleanExpr) SetLeft(left Expr) {
	b.Left = left
}

func (b *BooleanExpr) SetRight(right Expr) {
	b.Right = right
}

func (b *NotExpr) SetValue(left Expr) {
	b.value = left
}

func (b *NotExpr) GetValue() Expr {
	return b.value
}

// Interpret runs the BooleanExpr over a Vertex and VertexPatn to check for a match
func (b *BooleanExpr) Interpret(vertex *vertices.Vertex, pattern *VertexPatn) bool {

	left := b.GetLeft()
	right := b.GetRight()
	if l, ok := left.(NonTerminalExpr); ok {
		if r, ok := right.(NonTerminalExpr); ok {
			if b.Boolean == expressions.AND {
				return l.Interpret(vertex, pattern) && r.Interpret(vertex, pattern)
			}
			if b.Boolean == expressions.OR {
				return l.Interpret(vertex, pattern) || r.Interpret(vertex, pattern)
			}
			if b.Boolean == expressions.XOR {
				return expressions.XORExclusive(l.Interpret(vertex, pattern), r.Interpret(vertex, pattern))
			}
		}
	}

	return false
}

// BooleanPrecedence returns the precedence (order of importance)
func BooleanPrecedence(item BooleanExpr) int {
	if item.Boolean == expressions.AND {
		return 9
	} else if item.Boolean == expressions.OR {
		return 11
	} else if item.Boolean == expressions.XOR {
		return 10
	}
	return 20
}

// NotPrecedence returns the precedence (order of importance)
func NotPrecedence(item NotExpr) int {
	return 13
}
