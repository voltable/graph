package query

import "github.com/RossMerr/Caudex.Graph/graph"

// Where
func (q Query) Where(predicate func(*graph.Vertex) bool) Query {
	return Query{
		Iterate: func() Iterator {
			next := q.Iterate()

			return func() (item *graph.Vertex, ok bool) {
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
