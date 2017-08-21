package ast

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// Comparison operators
type Comparison int

const (
	// EQ equality
	EQ Comparison = iota // =
	// NEQ inequality
	NEQ // <>
	// LT less than
	LT // <
	// LTE less than or equal to
	LTE // <=
	// GT greater than
	GT // >
	// GTE greater than or equal to
	GTE // >=
	// IS_NULL used of IS NULL
	IS_NULL
	// IS_NOT_NULL used for IS NULL
	IS_NOT_NULL
)

// ComparisonExpr comparison expression
type ComparisonExpr struct {
	Comparison
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

	if b.Comparison == EQ {
		return resolve(b.GetX(), vertex, pattern) == resolve(b.GetY(), vertex, pattern)
	} else if b.Comparison == NEQ {
		return resolve(b.GetX(), vertex, pattern) != resolve(b.GetY(), vertex, pattern)
	} else {
		x := resolve(b.GetX(), vertex, pattern)
		y := resolve(b.GetY(), vertex, pattern)
		if b.Comparison == IS_NULL {
			return x == nil
		} else if b.Comparison == IS_NOT_NULL {
			return x != nil
		} else if b.Comparison == LT {
			switch i := x.(type) {
			case float64:
				return i < y.(float64)
			case float32:
				return i < y.(float32)
			case int64:
				return i < y.(int64)
			case int:
				return i < y.(int)
			case int16:
				return i < y.(int16)
			}
		} else if b.Comparison == LTE {
			switch i := x.(type) {
			case float64:
				return i <= y.(float64)
			case float32:
				return i <= y.(float32)
			case int64:
				return i <= y.(int64)
			case int:
				return i <= y.(int)
			case int16:
				return i <= y.(int16)
			}
		} else if b.Comparison == GT {
			switch i := x.(type) {
			case float64:
				return i > y.(float64)
			case float32:
				return i > y.(float32)
			case int64:
				return i > y.(int64)
			case int:
				return i > y.(int)
			case int16:
				return i > y.(int16)
			}
		} else if b.Comparison == GTE {
			switch i := x.(type) {
			case float64:
				return i >= y.(float64)
			case float32:
				return i >= y.(float32)
			case int64:
				return i >= y.(int64)
			case int:
				return i >= y.(int)
			case int16:
				return i >= y.(int16)
			}
		}
	}

	return false
}

// ComparisonPrecedence returns the precedence (order of importance)
func ComparisonPrecedence(item ComparisonExpr) int {
	if item.Comparison == EQ {
		return 8
	} else if item.Comparison == NEQ {
		return 8
	} else if item.Comparison == LT {
		return 7
	} else if item.Comparison == LTE {
		return 7
	} else if item.Comparison == GT {
		return 7
	} else if item.Comparison == GTE {
		return 7
	} else if item.Comparison == IS_NULL {
		return 7
	} else if item.Comparison == IS_NOT_NULL {
		return 7
	} else {
		return 20
	}
}
