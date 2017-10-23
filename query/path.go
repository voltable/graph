package query

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph/vertices"
)

type (
	// PredicateVertexPath is the Vertex implementation part of the QueryPath sequence
	PredicateVertexPath struct {
		PredicateVertex

		Variable string
	}

	// PredicateEdgePath is the Edge implementation part of the QueryPath sequence
	PredicateEdgePath struct {
		PredicateEdge

		Variable string
	}

	// IteratorFrontier is an alias for function to iterate over Frontier.
	IteratorFrontier func() (item *Frontier, ok Traverse)

	// PredicateVertex apply the predicate over the vertex
	PredicateVertex func(v *vertices.Vertex) (string, Traverse)

	// PredicateEdge apply the predicate over the edge
	PredicateEdge func(e *vertices.Edge, depth uint) (string, Traverse)
)

var (
	errPredicateVertexPath = errors.New("Expected PredicateVertexPath")
	errPredicateEdgePath   = errors.New("Expected PredicateEdgePath")
)
