package query

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// VertexPath is used to store data from the result of a Uniform Cost Search over vertexes.
//
// It only acts as one part of a Path from a walk in the graph you want to traverse acting on the Vertex.
// See EdgePath for walking over the Edge.
type VertexPath struct {
	Iterate  IteratorFrontier
	explored map[string]bool
	fetch    func(string) (*vertices.Vertex, error)
}

// NewVertexPath construts a new VertexPath
func NewVertexPath(i IteratorFrontier, f func(string) (*vertices.Vertex, error)) *VertexPath {
	return &VertexPath{explored: make(map[string]bool), fetch: f, Iterate: i}
}

// Node returns all Verteces matching the predicate
//
// The query is lazy only running on calling Iterate() from the EdgePath
func (t *VertexPath) Node(predicate PredicateVertex) *EdgePath {
	if predicate == nil {
		predicate = AllVertices()
	}

	return &EdgePath{
		explored: t.explored,
		fetch:    t.fetch,
		Iterate: func() (frontier *Frontier, ok bool) {
			for frontier, ok = t.Iterate(); ok; frontier, ok = t.Iterate() {
				vertices := frontier.peek()
				vertex := vertices[len(vertices)-1]
				t.explored[vertex.ID()] = true
				if predicate(vertex.Vertex) {
					return frontier, true
				}
			}
			return
		},
	}
}

// AllVertices matches all Vertexes.
func AllVertices() PredicateVertex {
	return func(v *vertices.Vertex) bool {
		return true
	}
}
