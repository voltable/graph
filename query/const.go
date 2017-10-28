package query

// Traverse is used to indicate the current state of the Traversal
type Traverse int

const (
	// Visiting is still traversing the graph
	Visiting Traverse = iota
	// Matching the edge's but not yet the vertex
	Matching
	// Matched the vertex
	Matched
	// Failed did not find a match in the traversal
	Failed
)
