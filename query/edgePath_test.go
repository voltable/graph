package query_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/expressions"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_MatchEdge(t *testing.T) {

	state := false
	vertex, _ := vertices.NewVertex()

	vertexDirection, _ := vertices.NewVertex()
	vertex.AddDirectedEdge(vertexDirection)

	frontier := query.Frontier{}
	fv := &query.FrontierVertex{Vertex: vertex}
	frontier = frontier.Append([]*query.FrontierVertex{fv}, 0)

	fetch := func(string) (*vertices.Vertex, error) {
		return vertex, nil
	}

	p := query.NewEdgePath(func() (item *query.Frontier, ok bool) {
		state = expressions.XORSwap(state)
		return &frontier, state
	}, fetch)

	matches := p.Relationship(func(v *vertices.Edge) (string, bool) {
		return "", true
	})

	matches.Iterate()
}
