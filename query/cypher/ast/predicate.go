package ast

import (
	"strings"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// ToPredicateVertex creates a PredicateVertex out of the VertexPatn
func ToPredicateVertex(patn *VertexPatn) query.PredicateVertex {
	return func(v *vertices.Vertex) bool {
		if strings.ToLower(patn.Label) != v.Label() {
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
	return func(v *vertices.Edge) bool {
		if strings.ToLower(patn.Body.Type) != v.RelationshipType() {
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
