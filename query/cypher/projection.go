package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

type projection struct {
}

func (p *projection) Transform(i query.IteratorFrontier, predicate ast.Expr) []interface{} {
	results := make([]interface{}, 0)
	for item, ok := i(); ok; item, ok = i() {
		for _, i := range item.OptimalPath() {
			results = append(results, i.Object)
		}

	}
	return results
}
