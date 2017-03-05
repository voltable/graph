package query

import "github.com/RossMerr/Caudex.Graph/vertices"

//Relationship
func (t Traversal) Relationship(predicate func(*vertices.Edge) bool) Traversal {
	return Traversal{
		iterate: func() Iterator {
			next := t.iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if e, is := item.(*vertices.Edge); is {
						if predicate(e) {
							return
						}
					}
				}
				return
			}
		},
	}
}
