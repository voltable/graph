package ast

import (
	"github.com/RossMerr/Caudex.Graph/comparisons"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// ComparisonExpr comparison expression
type ComparisonExpr struct {
	comparisons.Comparison
	Left  Expr // left operand
	Right Expr // right operand
}

func (ComparisonExpr) exprNode() {}

// GetLeft return value store in left side
func (b *ComparisonExpr) GetLeft() Expr {
	return b.Left
}

// GetRight return value store in right side
func (b *ComparisonExpr) GetRight() Expr {
	return b.Right
}

// SetLeft stores the Expr in left side
func (b *ComparisonExpr) SetLeft(left Expr) {
	b.Left = left
}

// SetRight stores the Expr in right side
func (b *ComparisonExpr) SetRight(right Expr) {
	b.Right = right
}

func resolve(expr Expr, vertex *vertices.Vertex, pattern *VertexPatn) interface{} {
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

	if b.Comparison == comparisons.EQ {
		return resolve(b.GetLeft(), vertex, pattern) == resolve(b.GetRight(), vertex, pattern)
	} else if b.Comparison == comparisons.NEQ {
		return resolve(b.GetLeft(), vertex, pattern) != resolve(b.GetRight(), vertex, pattern)
	} else {
		x := resolve(b.GetLeft(), vertex, pattern)
		y := resolve(b.GetRight(), vertex, pattern)

		return comparisons.Compare(b.Comparison, x, y)
	}
}

// ComparisonPrecedence returns the precedence (order of importance)
func ComparisonPrecedence(item ComparisonExpr) int {
	if item.Comparison == comparisons.EQ {
		return 8
	} else if item.Comparison == comparisons.NEQ {
		return 8
	} else if item.Comparison == comparisons.LT {
		return 7
	} else if item.Comparison == comparisons.LTE {
		return 7
	} else if item.Comparison == comparisons.GT {
		return 7
	} else if item.Comparison == comparisons.GTE {
		return 7
	} else if item.Comparison == comparisons.IS_NULL {
		return 7
	} else if item.Comparison == comparisons.IS_NOT_NULL {
		return 7
	} else {
		return 20
	}
}
