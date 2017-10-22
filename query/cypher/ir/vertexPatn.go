package ir

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

// ToPredicateVertexPath creates a PredicateVertexPath out of the VertexPatn
func (patn *VertexPatn) ToPredicateVertexPath() query.PredicateVertexPath {
	label := strings.ToLower(patn.Label)
	pvp := query.PredicateVertexPath{PredicateVertex: func(v *vertices.Vertex) (string, query.Traverse) {
		if label != v.Label() {
			return patn.Variable, query.Failed
		}

		for key, value := range patn.Properties {
			if v.Property(key) != value {
				return patn.Variable, query.Failed
			}
		}

		return patn.Variable, query.Visiting
	}, Variable: patn.Variable}

	return pvp
}
