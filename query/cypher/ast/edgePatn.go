package ast

import (
	"strings"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// Patn all pattern nodes implement the Patn interface.
type Patn interface {
	patnNode()
}

type EdgePatn struct {
	Relationship Digraph
	Body         *EdgeBodyStmt

	Vertex *VertexPatn
}

type EdgeBodyStmt struct {
	Variable      string
	Properties    map[string]interface{}
	Type          string
	LengthMinimum uint
	LengthMaximum uint
}

func (*EdgePatn) patnNode() {}

// ToPredicateEdge creates a PredicateEdge out of the EdgePatn
func (patn *EdgePatn) ToPredicateEdge() query.PredicateEdge {
	relationshipType := strings.ToLower(patn.Body.Type)
	return func(v *vertices.Edge) (string, bool) {
		if relationshipType != v.RelationshipType() {
			return patn.Body.Variable, false
		}

		for key, value := range patn.Body.Properties {
			if v.Property(key) != value {
				return patn.Body.Variable, false
			}
		}

		return patn.Body.Variable, true
	}
}
