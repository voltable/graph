package ast

import (
	"strings"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type VertexPatn struct {
	Variable   string
	Properties map[string]interface{}
	Label      string

	Edge *EdgePatn
}

func (*VertexPatn) patnNode() {}

// ToPredicateVertex creates a PredicateVertex out of the VertexPatn
func (patn *VertexPatn) ToPredicateVertex() query.PredicateVertex {
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
