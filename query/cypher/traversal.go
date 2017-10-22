package cypher

import (
	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/query"
)

type CypherTraversal interface {
	Travers(i enumerables.Iterator, path query.Path) (query.IteratorFrontier, error)
}
