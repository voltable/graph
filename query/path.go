package query

import (
	"container/list"
	"errors"

	"github.com/RossMerr/Caudex.Graph/keyvalue"
)

type (
	// Path is used to store data from the result of a Uniform Cost Search as part of the walk in the graph.
	Path struct {
		Iterate    IteratorFrontier
		explored   map[string]bool
		storage    keyvalue.Storage
		predicates *list.List
	}
)

var (
	errPredicateVertexPath = errors.New("Expected PredicatePath")
)
