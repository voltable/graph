package keyvalue

import "github.com/RossMerr/Caudex.Graph/uuid"

var (
	Vertex                 = []byte("v")
	Properties             = []byte("p")
	Relationship           = []byte("r")
	Relationshipproperties = []byte("k")
	// US unit separator can be used as delimiters to mark fields of data structures. If used for hierarchical levels, US is the lowest level (dividing plain-text data items)
	US = []byte(string('\u241F'))
)

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

// PredicatePath is the implementation part of the QueryPath sequence
type PredicatePath struct {
	Predicate
	Variable string
}

// Predicate apply the predicate over the key/value
type Predicate func(uuid uuid.UUID, depth int) (string, Traverse)
