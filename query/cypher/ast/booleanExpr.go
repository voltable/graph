package ast

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type Boolean int

const (
	// AND a bitwise or logical AND operation, such as (a And b).
	AND Boolean = iota
	// OR a bitwise or logical OR operation, such as (a Or b).
	OR
	// XOR a bitwise or logical XOR operation, such as (a Xor b).
	XOR

// NOT a bitwise complement or logical negation operation (Not a).
//NOT
)

type BooleanExpr struct {
	Boolean
	Left  Expr // left operand
	Right Expr // right operand
}

// NotExpr a bitwise complement or logical negation operation (Not a).
type NotExpr struct {
	Left Expr // left operand
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

func (b *NotExpr) SetLeft(left Expr) {
	b.Left = left
}

// Interpret runs the BooleanExpr over a Vertex and VertexPatn to check for a match
func (b *BooleanExpr) Interpret(vertex *vertices.Vertex, pattern *VertexPatn) bool {

	left := b.GetLeft()
	right := b.GetRight()
	if l, ok := left.(OperatorExpr); ok {
		if r, ok := right.(OperatorExpr); ok {
			if b.Boolean == AND {
				return l.Interpret(vertex, pattern) && r.Interpret(vertex, pattern)
			}
			if b.Boolean == OR {
				return l.Interpret(vertex, pattern) || r.Interpret(vertex, pattern)
			}
		}
	}

	return false
}

// BooleanPrecedence returns the precedence (order of importance)
func BooleanPrecedence(item BooleanExpr) int {
	if item.Boolean == AND {
		return 9
	} else if item.Boolean == OR {
		return 11
	} else if item.Boolean == XOR {
		return 10
	} else {
		return 20
	}
}

// NotPrecedence returns the precedence (order of importance)
func NotPrecedence(item NotExpr) int {
	return 13
}
