package graphs

// Where
func (q Query) Where(predicate func(*Vertex) bool) Query {
	return Query{
		Iterate: func() Iterator {
			next := q.Iterate()

			return func() (item *Vertex, ok bool) {
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
