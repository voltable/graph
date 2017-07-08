package query_test

import (
	"testing"

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

	it := func() (item interface{}, ok bool) {
		state = XOR(state)
		return frontier, state
	}

	p := query.NewVertexPath(func() query.Iterator {
		return it
	}, nil)

	matches := p.Node(func(v *vertices.Vertex) bool {
		if v.Label() != "foo" {
			t.Fatalf("Expected foo but was %s", v.Label())
		}
		return true
	})

	next := matches.Iterate()

	next()
}

// Swaps the boolean
func XOR(b bool) bool {
	if b == true {
		return false
	}

	return true
}
