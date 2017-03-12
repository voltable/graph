package query

import "github.com/RossMerr/Caudex.Graph/vertices"

// EdgePath represents the Edge part of a Path
type EdgePath struct {
	Iterate func() Iterator
	next    Path
	fetch   func(string) (*vertices.Vertex, error)
}

// Match returns all edges matching the predicate
func (t *EdgePath) Match(predicate func(*vertices.Edge) bool) *VertexPath {
	if predicate == nil {
		predicate = AllEdges()
	}

	return &VertexPath{
		Iterate: func() Iterator {
			next := t.Iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if e, is := item.(*vertices.Edge); is {
						if predicate(e) {
							if v, err := t.fetch(e.ID()); err != nil {
								return v, true
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
