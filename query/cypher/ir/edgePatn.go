package ir

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
	Variable string

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

// ToPredicateEdgePath creates a PredicateEdgePath out of the EdgePatn
func (patn *EdgePatn) ToPredicateEdgePath() query.PredicateEdgePath {
	relationshipType := strings.ToLower(patn.Body.Type)
	pvp := query.PredicateEdgePath{PredicateEdge: func(v *vertices.Edge, depth uint) (string, query.Traverse) {

		if depth < patn.Body.LengthMinimum {
			return patn.Body.Variable, query.Visiting
		}

		if depth > patn.Body.LengthMaximum {
			return patn.Body.Variable, query.Failed
		}

		if relationshipType != emptyString {
			if relationshipType != v.RelationshipType() {
				return patn.Body.Variable, query.Failed
			}
		}

		for key, value := range patn.Body.Properties {
			if v.Property(key) != value {
				return patn.Body.Variable, query.Failed
			}
		}

		return patn.Body.Variable, query.Visiting
	}, Variable: patn.Variable}

	return pvp
}
