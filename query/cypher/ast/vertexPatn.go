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

// ToPredicateVertexPath creates a PredicateVertexPath out of the VertexPatn
func (patn *VertexPatn) ToPredicateVertexPath() query.PredicateVertexPath {
	label := strings.ToLower(patn.Label)
	pvp := query.PredicateVertexPath{PredicateVertex: func(v *vertices.Vertex) (string, bool) {
		if label != v.Label() {
			return patn.Variable, false
		}

		for key, value := range patn.Properties {
			if v.Property(key) != value {
				return patn.Variable, false
			}
		}

		return patn.Variable, true
	}, Variable: patn.Variable}

	return pvp
}
