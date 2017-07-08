package query

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// type Predicate func(interface{}) bool

// type Query struct {
// 	Iterate        func() Iterator
// 	Explored       map[string]bool
// 	Fetch          func(string) (*vertices.Vertex, error)
// 	PrePredicate   []Predicate
// 	PostPredicatey []Predicate
// }

// EdgePath represents the Edge part of a Path
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
						path, frontier := frontier.pop()
						vertex := path.Vertices[len(path.Vertices)-1]
						for _, e := range vertex.Edges() {
							if _, ok := t.explored[e.ID()]; !ok {
								if predicate(e) {
									if v, err := t.fetch(e.ID()); err != nil {
										frontier = append(frontier, &Path{append(path.Vertices, v), path.Cost + e.Weight()})
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
			path, _ := frontier.pop()
			vertex := path.Vertices[len(path.Vertices)-1]
			slice = append(slice, vertex)

		}
	}
	return slice
}
