package query

import (
	"sort"

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
	fetch    func(string) (*vertices.Vertex, error)
}

// NewEdgePath constructs a new EdgePath
func NewEdgePath(i IteratorFrontier, f func(string) (*vertices.Vertex, error)) *EdgePath {
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
		Iterate: func() (frontier *Frontier, ok bool) {
			for frontier, ok = t.Iterate(); ok; frontier, ok = t.Iterate() {
				vertices, cost, frontier := frontier.Pop()
				vertex := vertices[len(vertices)-1]
				for _, e := range vertex.Edges() {
					if _, ok := t.explored[e.ID()]; !ok {
						if variable, p := predicate(e); p {
							if v, err := t.fetch(e.ID()); err == nil {
								fv := &FrontierVertex{Vertex: v, Variable: variable}
								frontier = frontier.Append(append(vertices, fv), cost+e.Weight)
								sort.Sort(frontier)
								return &frontier, true
							}
						}
					}
				}
			}
			return
		},
	}
}

// AllEdges matches all Edge.
func AllEdges() PredicateEdge {
	return func(v *vertices.Edge) (string, bool) {
		return emptyString, true
	}
}
