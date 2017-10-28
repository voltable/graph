package query

import (
	"sort"

	"github.com/RossMerr/Caudex.Graph/storage"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

var emptyString = ""

// NewEdgePath constructs a new edge Path
func NewEdgePath(i IteratorFrontier, s storage.Storage) *Path {
	return &Path{Iterate: i, storage: s}
}

// Relationship returns all edges matching the predicate
//
// The query is lazy only running on calling Iterate()
func (t *Path) Relationship(predicate PredicateEdge) *Path {
	if predicate == nil {
		predicate = AllEdges()
	}

	return &Path{
		explored: t.explored,
		storage:  t.storage,
		Iterate: func() (frontier *Frontier, ok Traverse) {
			for frontier, ok = t.Iterate(); ok != Failed; frontier, ok = t.Iterate() {
				vertices, cost := frontier.Pop()
				depth := len(vertices)
				vertex := vertices[depth-1]
				for _, e := range vertex.Edges() {
					if _, ok := t.explored[e.ID()]; !ok {
						if variable, p := predicate(e, uint(depth)); p != Failed {
							if v, err := t.storage.Fetch(e.ID()); err == nil {
								fv := &FrontierVertex{Vertex: v, Variable: variable}
								frontier.Append(append(vertices, fv), cost+e.Weight, p)
							}
						}
					}
				}
				sort.Sort(frontier)
				return frontier, Visiting
			}
			return
		},
	}
}

// AllEdges matches all Edge.
func AllEdges() PredicateEdge {
	return func(v *vertices.Edge, depth uint) (string, Traverse) {
		return emptyString, Matched
	}
}
