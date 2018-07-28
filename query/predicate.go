package query

import "github.com/RossMerr/Caudex.Graph/uuid"

// Predicate apply the predicate over the key/value
//
// string Variable
// Traverse Traverse
// float64 Weight
type Predicate func(from, to *uuid.UUID, depth int) (string, Traverse, float64)

// Traverse is used to indicate the current state of the Traversal
type Traverse int

const (
	// Visiting is traversing the graph and not matching any part of the edge or vertex
	Visiting Traverse = iota
	// Matching the edge's but not yet the vertex so mighe be traversing the edge and vertex
	Matching
	// Matched the vertex
	Matched
	// Failed did not find a match in the traversal
	Failed
)
