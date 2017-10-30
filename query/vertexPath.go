package query

import (
	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/storage"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// NewVertexPath construts a new vertex Path
func NewVertexPath(i enumerables.Iterator, s storage.Storage, variable string) *Path {
	return &Path{explored: make(map[string]bool), storage: s, Iterate: toFontier(i, variable)}
}

// Node returns all Verteces matching the predicate
//
// The query is lazy only running on calling Iterate()
func (t *Path) Node(predicate PredicateVertex) *Path {
	if predicate == nil {
		predicate = AllVertices()
	}

	return &Path{
		explored: t.explored,
		storage:  t.storage,
		Iterate: func() (frontier *Frontier, ok Traverse) {
			for frontier, ok = t.Iterate(); ok != Failed; frontier, ok = t.Iterate() {
				if frontier.Len() > 0 {
					vertices, _ := frontier.OptimalPath()
					vertex := vertices[len(vertices)-1]
					t.explored[vertex.ID()] = true
					if variable, p := predicate(vertex.Vertex); p != Failed {
						vertex.Variable = variable
						return frontier, p
					}
				}
			}
			return
		},
	}
}

// AllVertices matches all Vertexes.
func AllVertices() PredicateVertex {
	return func(v *vertices.Vertex) (string, Traverse) {
		return "", Visiting
	}
}

func toFontier(i enumerables.Iterator, variable string) IteratorFrontier {
	return func() (*Frontier, Traverse) {
		for item, ok := i(); ok; item, ok = i() {
			if v, is := item.(*vertices.Vertex); is {
				f := NewFrontier(v, variable)
				return &f, Visiting
			}
		}
		return nil, Failed
	}
}
