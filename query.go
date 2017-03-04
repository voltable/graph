package graphs

type (
	// Iterator is an alias for function to iterate over data.
	Iterator func() (item *Vertex, ok bool)

	// Query is the type returned from query functions.
	Query struct {
		Iterate func() Iterator
	}
)
