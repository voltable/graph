package cypher

import (
	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/query"
)

type CypherTraversal interface {
	//Travers(i enumerables.Iterator, path *list.List) (query.IteratorFrontier, error)

	SearchPlan(iterator enumerables.Iterator, predicates []interface{}) (iteratorFrontier query.IteratorFrontier, err error)
}
