package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

type QueryBuilder interface {
	Predicate([]ast.Patn) []query.Predicate
}
