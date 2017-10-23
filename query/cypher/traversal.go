package cypher

import (
	"container/list"

	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/query"
)

type CypherTraversal interface {
	Travers(i enumerables.Iterator, path *list.List) (query.IteratorFrontier, error)
}
