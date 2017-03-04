package query

import "github.com/RossMerr/Caudex.Graph/graph/vertices"

type (
	// Iterator is an alias for function to iterate over data.
	Iterator func() (item *vertices.Vertex, ok bool)

	// Query is the type returned from query functions.
	Query struct {
		Iterate func() Iterator
	}
)
