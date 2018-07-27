package cypher

import (
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

type Projection struct {
}

func (p *Projection) Transform(i query.IteratorFrontier, maps []*ast.ProjectionMapStmt) []interface{} {
	results := make([]interface{}, 0)
	for item, ok := i(); ok; item, ok = i() {
		for _, part := range item.OptimalPath() {
			for _, m := range maps {
				results = append(results, m.Interpret(part.Variable, part.KeyValue)...)
			}
		}

	}
	return results
}
