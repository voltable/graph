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

	iterator := func() (item interface{}, ok bool) {
		state = expressions.XORSwap(state)
		return vertex, state
	}

	p := query.NewVertexPath(iterator, nil)

	matches := p.Node(func(v *vertices.Vertex) bool {
		if v.Label() != "foo" {
			t.Fatalf("Expected foo but was %s", v.Label())
		}
		return true
	})

	matches.Iterate()

}
