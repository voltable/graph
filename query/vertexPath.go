package query

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// VertexPath represents the Vertex part of a Path
type VertexPath struct {
	Iterate func() Iterator
	next    Path
	fetch   func(string) (*vertices.Vertex, error)
}

// Match returns all Verteces matching the predicate
func (t VertexPath) Match(predicate func(v *vertices.Vertex) bool) EdgePath {
	return EdgePath{
		Iterate: func() Iterator {
			next := t.Iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if v, is := item.(*vertices.Vertex); is {
						if predicate(v) {
							for _, e := range v.Edges() {
								return e, true
							}
						}
					}
				}
				return
			}
		},
	}
}
