package query

import (
	"sort"

	"github.com/RossMerr/Caudex.Graph/vertices"
)

// VertexPath represents the Vertex part of a Path
type VertexPath struct {
	Iterate  func() Iterator
	explored map[string]bool
	fetch    func(string) (*vertices.Vertex, error)
}

// NewVertexPath construts a new VertexPath
func NewVertexPath(i func() Iterator, f func(string) (*vertices.Vertex, error)) *VertexPath {
	return &VertexPath{explored: make(map[string]bool), fetch: f, Iterate: i}
}

// Node returns all Verteces matching the predicate
func (t *VertexPath) Node(predicate PredicateVertex) *EdgePath {
	if predicate == nil {
		predicate = AllVertices()
	}

	return &EdgePath{
		explored: t.explored,
		fetch:    t.fetch,
		Iterate: func() Iterator {
			next := t.Iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if frontier, is := item.(Frontier); is {
						sort.Sort(frontier)
						vertices := frontier.peek()
						vertex := vertices[len(vertices)-1]
						t.explored[vertex.ID()] = true
						if predicate(vertex) {
							return frontier, true
						}
					}
				}
				return
			}
		},
	}
}

// AllVertices matches all Vertexes.
func AllVertices() PredicateVertex {
	return func(v *vertices.Vertex) bool {
		return true
	}
}
