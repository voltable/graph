package query

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// EdgePath represents the Edge part of a Path
type EdgePath struct {
	Iterate  func() Iterator
	Explored map[string]bool
	Fetch    func(string) (*vertices.Vertex, error)
}

// Match returns all edges matching the predicate
func (t *EdgePath) Match(predicate func(*vertices.Edge) bool) *VertexPath {
	if predicate == nil {
		predicate = AllEdges()
	}

	return &VertexPath{
		Explored: t.Explored,
		Fetch:    t.Fetch,
		Iterate: func() Iterator {
			next := t.Iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if frontier, is := item.(Frontier); is {
						path, frontier := frontier.pop()
						vertex := path.Vertices[len(path.Vertices)-1]
						for _, e := range vertex.Edges() {
							if _, ok := t.Explored[e.ID()]; !ok {
								if predicate(e) {
									if v, err := t.Fetch(e.ID()); err != nil {
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

func (t *EdgePath) MatchAll() *VertexPath {
	return t.Match(nil)
}

// AllEdges matches all Edge.
func AllEdges() func(v *vertices.Edge) bool {
	return func(v *vertices.Edge) bool {
		return true
	}
}
