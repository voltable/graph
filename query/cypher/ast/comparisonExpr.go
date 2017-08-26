package ast

import (
	"github.com/RossMerr/Caudex.Graph/comparisons"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// ComparisonExpr comparison expression
type ComparisonExpr struct {
	comparisons.Comparison
	X Expr // left operand
	Y Expr // right operand
}

func (ComparisonExpr) exprNode() {}

// GetX return value store in X, left side
func (b *ComparisonExpr) GetX() Expr {
	return b.X
}

// GetY return value store in Y, right side
func (b *ComparisonExpr) GetY() Expr {
	return b.Y
}

// SetX stores the Expr in X, left side
func (b *ComparisonExpr) SetX(x Expr) {
	b.X = x
}

// SetY stores the Expr in Y, right side
func (b *ComparisonExpr) SetY(y Expr) {
	b.Y = y
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

// Evaluate runs the ComparisonExpr over a Vertex and VertexPatn to check for a match
//
// The ComparisonExpr comes from building the AST so it is part of the WHERE clause
//     WHERE n.age < 30
// The VertexPatn is part of the a MATCH statment within the query
//     MATCH (n:Person)
// Finally the Vertex is the vertex you want to run the Evaluate over to check for a match
func (b *ComparisonExpr) Evaluate(vertex *vertices.Vertex, pattern *VertexPatn) bool {

	if b.Comparison == comparisons.EQ {
		return resolve(b.GetX(), vertex, pattern) == resolve(b.GetY(), vertex, pattern)
	} else if b.Comparison == comparisons.NEQ {
		return resolve(b.GetX(), vertex, pattern) != resolve(b.GetY(), vertex, pattern)
	} else {
		x := resolve(b.GetX(), vertex, pattern)
		y := resolve(b.GetY(), vertex, pattern)

		return comparisons.Compare(b.Comparison, x, y)
	}

	return false
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
