package query

import "github.com/RossMerr/Caudex.Graph/vertices"

// MatchEdge returns all edges matching the predicate
func (t EdgePath) MatchEdge(predicate func(*vertices.Edge) bool) VertexPath {
	return VertexPath{
		Iterate: func() Iterator {
			next := t.Iterate()
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
