package ast

import (
	"github.com/RossMerr/Caudex.Graph/expressions"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// ComparisonExpr comparison expression
type ComparisonExpr struct {
	expressions.Comparison
	left  InterpretExpr // left operand
	right InterpretExpr // right operand
}

var _ NonTerminalExpr = (*ComparisonExpr)(nil)

func (ComparisonExpr) exprNode()      {}
func (ComparisonExpr) interpretNode() {}

// NewComparisonExpr creates a ComparisonExpr
func NewComparisonExpr(comparison expressions.Comparison, left InterpretExpr, right InterpretExpr) *ComparisonExpr {
	return &ComparisonExpr{Comparison: comparison, left: left, right: right}
}

// GetLeft return value store in left side
func (b *ComparisonExpr) GetLeft() InterpretExpr {
	return b.left
}

// GetRight return value store in right side
func (b *ComparisonExpr) GetRight() InterpretExpr {
	return b.right
}

// SetLeft stores the Expr in left side
func (b *ComparisonExpr) SetLeft(left InterpretExpr) {
	b.left = left
}

// SetRight stores the Expr in right side
func (b *ComparisonExpr) SetRight(right InterpretExpr) {
	b.right = right
}

func resolve(expr InterpretExpr, vertex *vertices.Vertex, pattern *VertexPatn) interface{} {
	if prop, ok := expr.(PropertyStmt); ok {
		if prop.Variable == pattern.Variable {
			return vertex.Property(prop.Value)
		}
		return false
	} else if prop, ok := expr.(Ident); ok {
		return prop.Data
	}

	return nil
}

// Interpret runs the ComparisonExpr over a Vertex and VertexPatn to check for a match
//
// The ComparisonExpr comes from building the AST so it is part of the WHERE clause
//     WHERE n.age < 30
// The VertexPatn is part of the a MATCH statement within the query
//     MATCH (n:Person)
// Finally the Vertex is the vertex you want to run the Evaluate over to check for a match
func (b *ComparisonExpr) Interpret(vertex *vertices.Vertex, pattern *VertexPatn) bool {

	left := b.GetLeft()
	right := b.GetRight()

	if b.Comparison == expressions.EQ {
		return resolve(left, vertex, pattern) == resolve(right, vertex, pattern)
	} else if b.Comparison == expressions.NEQ {
		return resolve(left, vertex, pattern) != resolve(right, vertex, pattern)
	} else {
		x := resolve(left, vertex, pattern)
		y := resolve(right, vertex, pattern)

		return expressions.Compare(b.Comparison, x, y)
	}
}

// ComparisonPrecedence returns the precedence (order of importance)
func ComparisonPrecedence(item ComparisonExpr) int {
	if item.Comparison == expressions.EQ {
		return 8
	} else if item.Comparison == expressions.NEQ {
		return 8
	} else if item.Comparison == expressions.LT {
		return 7
	} else if item.Comparison == expressions.LTE {
		return 7
	} else if item.Comparison == expressions.GT {
		return 7
	} else if item.Comparison == expressions.GTE {
		return 7
	} else if item.Comparison == expressions.IS_NULL {
		return 7
	} else if item.Comparison == expressions.IS_NOT_NULL {
		return 7
	} else {
		return 20
	}
}
