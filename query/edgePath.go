package query

import (
	"sort"

	"github.com/RossMerr/Caudex.Graph/storage"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

var emptyString = ""

// EdgePath is used to store data from the result of a Uniform Cost Search over edges.
//
// It only acts as one part of a Path from a walk in the graph you want to traverse acting on the edge.
// See VertexPath for walking over the Vertices.
type EdgePath struct {
	Iterate  IteratorFrontier
	explored map[string]bool
	storage  storage.Storage
}

// NewEdgePath constructs a new EdgePath
func NewEdgePath(i IteratorFrontier, s storage.Storage) *EdgePath {
	return &EdgePath{Iterate: i, storage: s}
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
		storage:  t.storage,
		Iterate: func() (frontier *Frontier, ok Traverse) {
			for frontier, ok = t.Iterate(); ok != Failed; frontier, ok = t.Iterate() {
				// if ok == Matched {
				// 	return
				// }
				vertices, cost, frontier := frontier.Pop()
				depth := len(vertices)
				vertex := vertices[depth-1]
				for _, e := range vertex.Edges() {
					if _, ok := t.explored[e.ID()]; !ok {
						if variable, p := predicate(e, uint(depth)); p != Failed {
							if v, err := t.storage.Fetch(e.ID()); err == nil {
								fv := &FrontierVertex{Vertex: v, Variable: variable}
								frontier = frontier.Append(append(vertices, fv), cost+e.Weight)
							}
						}
					}
				}
				sort.Sort(frontier)
				return &frontier, Visiting
			}
			return

			// for frontier, ok = t.Iterate(); ok == Visiting; frontier, ok = t.Iterate() {
			// 	vertices, cost, frontier := frontier.Pop()
			// 	depth := len(vertices)
			// 	vertex := vertices[depth-1]
			// 	for _, e := range vertex.Edges() {
			// 		if _, ok := t.explored[e.ID()]; !ok {
			// 			if variable, p := predicate(e, uint(depth)); p == Visiting || p == Matched {
			// 				if v, err := t.storage.Fetch(e.ID()); err == nil {
			// 					fv := &FrontierVertex{Vertex: v, Variable: variable}
			// 					frontier = frontier.Append(append(vertices, fv), cost+e.Weight)
			// 					sort.Sort(frontier)
			// 					return &frontier, Visiting
			// 				}
			// 			}
			// 		}
			// 	}
			// }
			// return
		},
	}
}

// AllEdges matches all Edge.
func AllEdges() PredicateEdge {
	return func(v *vertices.Edge, depth uint) (string, Traverse) {
		return emptyString, Matched
	}
}
