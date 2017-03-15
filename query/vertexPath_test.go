package query

import "testing"
import "github.com/RossMerr/Caudex.Graph/vertices"

func Test_MatchVertex(t *testing.T) {

	state := false
	vertex, _ := vertices.NewVertex()
	vertex.SetLabel("foo")

	vertexDirection, _ := vertices.NewVertex()
	vertex.AddDirectedEdge(vertexDirection)

	frontier := Frontier{&Path{[]*vertices.Vertex{vertex}, 0}}

	it := func() (item interface{}, ok bool) {
		state = XOR(state)
		return frontier, state
	}

	p := VertexPath{Iterate: func() Iterator {
		return it
	}, Explored: make(map[string]bool)}

	matches := p.Match(func(v *vertices.Vertex) bool {
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
