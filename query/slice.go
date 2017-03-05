package query

import (
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// ToSlice iterates over a collection and saves the results in the slice
func (q Query) ToSlice(slice []*vertices.Vertex) {
	next := q.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		slice = append(slice, item)
	}
}
