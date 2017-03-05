package query

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// Select
func (t Traversal) Select(predicate func(*vertices.Vertex) bool) Traversal {
	return Traversal{
		iterate: func() Iterator {
			next := t.iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if v, is := item.(*vertices.Vertex); is {
						if predicate(v) {
							return
						}
					}
				}
				return
			}
		},
	}
}

