package query

import "github.com/RossMerr/Caudex.Graph/graph/vertices"

// Where
func (q Query) Where(predicate func(*vertices.Vertex) bool) Query {
	return Query{
		Iterate: func() Iterator {
			next := q.Iterate()

			return func() (item *vertices.Vertex, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if predicate(item) {
						return
					}
				}

				return
			}
		},
	}
}
