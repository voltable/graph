package graphs

// Query is the type returned from query functions.
type Query struct {
	Iterate func() Iterator
}
