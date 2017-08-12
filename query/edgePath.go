package query

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// EdgePath is used to store data from the result of a Uniform Cost Search over edges.
//
// It only acts as one part of a Path from a walk in the graph you want to traverse acting on the edge part.
// See VertexPath for walking over the Vertices.
type EdgePath struct {
	Iterate  func() Iterator
	explored map[string]bool
	fetch    func(string) (*vertices.Vertex, error)
}

// NewEdgePath constructs a new EdgePath
func NewEdgePath(i func() Iterator, f func(string) (*vertices.Vertex, error)) *EdgePath {
	return &EdgePath{Iterate: i, fetch: f, explored: make(map[string]bool)}
}

// Relationship returns all edges matching the predicate
//
// The query is lazy only running on calling Iterate() from the VertexPath
func (t *EdgePath) Relationship(predicate PredicateEdge) *VertexPath {
	if predicate == nil {
		predicate = AllEdges()
	}

	return &VertexPath{
		explored: t.explored,
		fetch:    t.fetch,
		Iterate: func() Iterator {
			next := t.Iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if frontier, is := item.(Frontier); is {
						vertices, cost, frontier := frontier.pop()
						vertex := vertices[len(vertices)-1]
						for _, e := range vertex.Edges() {
							if _, ok := t.explored[e.ID()]; !ok {
								if predicate(e) {
									if v, err := t.fetch(e.ID()); err == nil {
										frontier = frontier.Append(append(vertices, v), cost+e.Weight())
										return frontier, true
									}
								}
							}
						}
					}
				}
				return
			}
		},
	}
}

// AllEdges matches all Edge.
func AllEdges() PredicateEdge {
	return func(v *vertices.Edge) bool {
		return true
	}
}

// ToSlice returns the final matching Vertexs of the query to a slice
func (t *EdgePath) ToSlice() []*vertices.Vertex {

	slice := []*vertices.Vertex{}
	next := t.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		if frontier, is := item.(Frontier); is {
			vertices, _, _ := frontier.pop()
			vertex := vertices[len(vertices)-1]
			slice = append(slice, vertex)

		}
	}
	return slice
}
