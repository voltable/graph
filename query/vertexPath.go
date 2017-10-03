package query

import (
	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/storage"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// VertexPath is used to store data from the result of a Uniform Cost Search over vertexes.
//
// It only acts as one part of a Path from a walk in the graph you want to traverse acting on the Vertex.
// See EdgePath for walking over the Edge.
type VertexPath struct {
	Iterate  IteratorFrontier
	explored map[string]bool
	storage  storage.Storage
}

// NewVertexPath construts a new VertexPath
func NewVertexPath(i enumerables.Iterator, s storage.Storage, variable string) *VertexPath {
	return &VertexPath{explored: make(map[string]bool), storage: s, Iterate: toFontier(i, variable)}
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
		storage:  t.storage,
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

func toFontier(i enumerables.Iterator, variable string) IteratorFrontier {
	return func() (*Frontier, bool) {
		for item, ok := i(); ok; item, ok = i() {
			if v, is := item.(*vertices.Vertex); is {
				f := NewFrontier(v, variable)
				return &f, true
			}
		}
		return nil, false
	}
}
