package ast

import (
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// BooleanExpr boolean expression
type BooleanExpr struct {
	expressions.Boolean
	left  InterpretExpr // left operand
	right InterpretExpr // right operand
}

var _ NonTerminalExpr = (*BooleanExpr)(nil)

func (BooleanExpr) exprNode()      {}
func (BooleanExpr) interpretNode() {}

// NewBooleanExpr creates a BooleanExpr
func NewBooleanExpr(boolean expressions.Boolean, left InterpretExpr, right InterpretExpr) *BooleanExpr {
	return &BooleanExpr{Boolean: boolean, left: left, right: right}
}

// GetLeft return value store in left side
func (b *BooleanExpr) GetLeft() InterpretExpr {
	return b.left
}

// GetRight return value store in right side
func (b *BooleanExpr) GetRight() InterpretExpr {
	return b.right
}

// SetLeft stores the Expr in left side
func (b *BooleanExpr) SetLeft(left InterpretExpr) {
	b.left = left
}

// SetRight stores the Expr in right side
func (b *BooleanExpr) SetRight(right InterpretExpr) {
	b.right = right
}

// Interpret runs the BooleanExpr over a Vertex to check for a match
func (b *BooleanExpr) Interpret(vertex *vertices.Vertex) interface{} {

	left := b.GetLeft()
	right := b.GetRight()

	x := left.Interpret(vertex).(bool)
	y := right.Interpret(vertex).(bool)
	if b.Boolean == expressions.AND {
		return x && y
	}
	if b.Boolean == expressions.OR {
		return x || y
	}
	if b.Boolean == expressions.XOR {
		return expressions.XORExclusive(x, y)
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
