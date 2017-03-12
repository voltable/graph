package query

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_MatchEdge(t *testing.T) {

	state := false
	vertex, _ := vertices.NewVertex()

	vertexDirection, _ := vertices.NewVertex()
	edge, _ := vertex.AddDirectedEdge(vertexDirection)
	edge.SetLabel("foo")

	it := func() (item interface{}, ok bool) {
		state = XOR(state)
		return edge, state
	}

	fetch := func(string) (*vertices.Vertex, error) {
		return vertex, nil
	}

	p := EdgePath{Iterate: func() Iterator {
		return it
	}, fetch: fetch}

	matches := p.Match(func(v *vertices.Edge) bool {
		if v.Label() != "foo" {
			t.Fatalf("Expected foo but was %s", edge.Label())
		}
		return true
	})

	matches.Iterate()()
}
