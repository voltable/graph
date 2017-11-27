package query

import (
	"container/list"
	"errors"

	"github.com/RossMerr/Caudex.Graph/storage"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

type (
	// Path is used to store data from the result of a Uniform Cost Search as part of the walk in the graph.
	Path struct {
		Iterate    IteratorFrontier
		explored   map[string]bool
		storage    storage.Storage
		predicates *list.List
	}

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
	IteratorFrontier func() (item *Frontier, ok bool)

	// PredicateVertex apply the predicate over the vertex
	PredicateVertex func(v *vertices.Vertex) (string, Traverse)

	// PredicateEdge apply the predicate over the edge
	PredicateEdge func(e *vertices.Edge, depth uint) (string, Traverse)
)

var (
	errPredicateVertexPath = errors.New("Expected PredicateVertexPath")
	errPredicateEdgePath   = errors.New("Expected PredicateEdgePath")
)
