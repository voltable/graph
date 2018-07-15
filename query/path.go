package query

import (
	"container/list"
	"errors"

	"github.com/RossMerr/Caudex.Graph/keyvalue"

	"github.com/RossMerr/Caudex.Graph"
)

type (
	// Path is used to store data from the result of a Uniform Cost Search as part of the walk in the graph.
	Path struct {
		Iterate    IteratorFrontier
		explored   map[string]bool
		storage    graph.Storage
		predicates *list.List
	}

	// PredicatePath is the implementation part of the QueryPath sequence
	PredicatePath struct {
		Predicate
		Variable string
	}

	// IteratorFrontier is an alias for function to iterate over Frontier.
	IteratorFrontier func() (item *Frontier, ok bool)

	// Predicate apply the predicate over the key/value
	Predicate func(kv *keyvalue.KeyValue) (string, Traverse)
)

var (
	errPredicateVertexPath = errors.New("Expected PredicatePath")
)
