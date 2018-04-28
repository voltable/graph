package triplestore

import (
	"strings"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"
)

type VertexPatn struct {
	Variable   string
	Properties map[string]interface{}
	Label      string

	Edge *EdgePatn
}

func (*VertexPatn) patnNode() {}

// ToPredicateVertexPath creates a PredicateVertexPath out of the VertexPatn
func (patn *VertexPatn) ToPredicateVertexPath() *query.PredicateVertexPath {
	label := strings.ToLower(patn.Label)
	pvp := query.PredicateVertexPath{PredicateVertex: func(v *graph.Vertex) (string, query.Traverse) {
		if label != v.Label() {
			return patn.Variable, query.Failed
		}

		for key, value := range patn.Properties {
			if v.Property(key) != value {
				return patn.Variable, query.Failed
			}
		}

		return patn.Variable, query.Matched

	}, Variable: patn.Variable}

	return &pvp
}
