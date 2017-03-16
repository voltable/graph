package query

import (
	"sort"

	"github.com/RossMerr/Caudex.Graph/vertices"
)

// VertexPath represents the Vertex part of a Path
type VertexPath struct {
	Iterate  func() Iterator
	Explored map[string]bool
	Fetch    func(string) (*vertices.Vertex, error)
}

func NewVertexPath(i func() Iterator, f func(string) (*vertices.Vertex, error)) *VertexPath {
	return &VertexPath{Explored: make(map[string]bool), Fetch: f, Iterate: i}
}

// Node returns all Verteces matching the predicate
func (t *VertexPath) Node(predicate PredicateVertex) *VertexToRelationshipPath {
	if predicate == nil {
		predicate = AllVertices()
	}

	return &VertexToRelationshipPath{
		Explored: t.Explored,
		Fetch:    t.Fetch,
		Iterate: func() Iterator {
			next := t.Iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if frontier, is := item.(Frontier); is {
						sort.Sort(frontier)
						path := frontier.peek()
						vertex := path.Vertices[len(path.Vertices)-1]
						t.Explored[vertex.ID()] = true
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
