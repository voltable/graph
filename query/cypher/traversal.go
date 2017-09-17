package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query"
)

type CypherTraversal interface {
	Travers(i query.IteratorFrontier, path query.Path) query.IteratorFrontier
}
