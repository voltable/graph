package ast

import (
	"fmt"
	"strings"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// ToPredicateVertex creates a PredicateVertex out of the VertexPatn
func ToPredicateVertex(patn *VertexPatn) query.PredicateVertex {
	label := strings.ToLower(patn.Label)
	return func(v *vertices.Vertex) bool {
		if label != v.Label() {
			return false
		}

		for key, value := range patn.Properties {
			if v.Property(key) != value {
				return false
			}
		}

		return true
	}
}

// ToPredicateEdge creates a PredicateEdge out of the EdgePatn
func ToPredicateEdge(patn *EdgePatn) query.PredicateEdge {
	relationshipType := strings.ToLower(patn.Body.Type)
	return func(v *vertices.Edge) bool {
		if relationshipType != v.RelationshipType() {
			return false
		}

		for key, value := range patn.Body.Properties {
			if v.Property(key) != value {
				return false
			}
		}

		return true
	}
}

// ToPredicateExpression creates a PredicateExpression out of the Expr
func ToPredicateExpression(item Expr) (query.PredicateExpression, error) {

	// if b, ok := item.(*BooleanExpr); ok {
	// 	return BooleanPrecedence(*b)
	// } else if b, ok := item.(*NotExpr); ok {
	// 	return NotPrecedence(*b)
	// } else if b, ok := item.(*ComparisonExpr); ok {
	// 	return ComparisonPrecedence(*b)
	// } else if b, ok := item.(MathematicalExpr); ok {
	// 	return MathPrecedence(b)
	// } else if b, ok := item.(ParenthesesExpr); ok {
	// 	return ParenthesesPrecedence(b)
	// }

	// if b, ok := item.(*ComparisonExpr); ok {
	// 	return func(path *query.Path) bool {
	// 		strings.Compare(b.)
	// 		b.Comparison.
	// 		return false
	// 	}, nil
	// }
	return func(path *query.Path) bool {

		return false
	}, fmt.Errorf("")
}
