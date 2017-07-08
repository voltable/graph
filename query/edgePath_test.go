package query_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_MatchEdge(t *testing.T) {

	state := false
	vertex, _ := vertices.NewVertex()

	vertexDirection, _ := vertices.NewVertex()
	vertex.AddDirectedEdge(vertexDirection)
	//edge.SetLabel("foo")

	frontier := query.Frontier{&query.Path{[]*vertices.Vertex{vertex}, 0}}

	it := func() (item interface{}, ok bool) {
		state = XOR(state)
		return frontier, state
	}

	fetch := func(string) (*vertices.Vertex, error) {
		return vertex, nil
	}

	p := query.NewEdgePath(func() query.Iterator {
		return it
	}, fetch)

	matches := p.Relationship(func(v *vertices.Edge) bool {
		// if v.Label() != "foo" {
		// 	t.Fatalf("Expected foo but was %s", v.Label())
		// }
		return true
	})

	matches.Iterate()()
}
