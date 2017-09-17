package query_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/expressions"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_MatchVertex(t *testing.T) {

	state := false
	vertex, _ := vertices.NewVertex()
	vertex.SetLabel("foo")

	vertexDirection, _ := vertices.NewVertex()
	vertex.AddDirectedEdge(vertexDirection)

	frontier := query.Frontier{}
	frontier = frontier.Append([]*vertices.Vertex{vertex}, 0)

	p := query.NewVertexPath(func() (item *query.Frontier, ok bool) {
		state = expressions.XORSwap(state)
		return &frontier, state
	}, nil)

	matches := p.Node(func(v *vertices.Vertex) bool {
		if v.Label() != "foo" {
			t.Fatalf("Expected foo but was %s", v.Label())
		}
		return true
	})

	matches.Iterate()

}
