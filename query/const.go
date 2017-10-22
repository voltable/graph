package query

// Traverse is used to indicate the current state of the Traversal
type Traverse int

const (
	// Visiting is still traversing the graph
	Visiting Traverse = iota
	// Failed  did not find a match in the traversal
	Failed
	// Matched in the traversal
	Matched
)
