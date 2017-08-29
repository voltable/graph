package ast

import (
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// BooleanExpr boolean expression
type BooleanExpr struct {
	expressions.Boolean
	left  Expr // left operand
	right Expr // right operand
}

var _ NonTerminalExpr = (*BooleanExpr)(nil)

func (BooleanExpr) exprNode() {}

// NewBooleanExpr creates a BooleanExpr
func NewBooleanExpr(boolean expressions.Boolean, left Expr, right Expr) *BooleanExpr {
	return &BooleanExpr{Boolean: boolean, left: left, right: right}
}

// GetLeft return value store in left side
func (b *BooleanExpr) GetLeft() Expr {
	return b.left
}

// GetRight return value store in right side
func (b *BooleanExpr) GetRight() Expr {
	return b.right
}

// SetLeft stores the Expr in left side
func (b *BooleanExpr) SetLeft(left Expr) {
	b.left = left
}

// SetRight stores the Expr in right side
func (b *BooleanExpr) SetRight(right Expr) {
	b.right = right
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
