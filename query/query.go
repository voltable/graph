package query

type (
	// Iterator is an alias for function to iterate over data.
	Iterator func() (item interface{}, ok bool)

	// MatchPattern is used for describing the structure of the pattern searched for
	MatchPattern struct {
		match map[string]Traversal
	}

	Traversal struct {
		iterate func() Iterator
	}
)
