package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query"
)

type CypherTraversal interface {
	Travers(i func() query.IteratorFrontier, path query.Path) func() query.IteratorFrontier
}
