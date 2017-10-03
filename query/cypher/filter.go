package cypher

import (
	"github.com/RossMerr/Caudex.Graph/enumerables"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

// CypherFilter the interface of the filter struct
type CypherFilter interface {
	Filter(query.IteratorFrontier, ast.Expr) enumerables.Iterator
}

// Filter is use to filter the traveral results over the where expression in the AST
type Filter struct {
}

// NewFilter create a new filter
func NewFilter() *Filter {
	return &Filter{}
}

// Filter a IteratorFrontier so that all results pass the Where Expression in the AST
func (qe Filter) Filter(i query.IteratorFrontier, predicate ast.Expr) enumerables.Iterator {
	return func() (interface{}, bool) {
		for frontier, ok := i(); ok; frontier, ok = i() {
			if frontier.Len() > 0 {
				// We only need the first array of vertices from the frontier as the rest aren't the the optimal path
				vertices, _, _ := frontier.Pop()
				for _, v := range vertices {
					if predicate != nil {
						if qe.ExpressionEvaluator(predicate, v.Variable, v.Vertex) {
							return v.Vertex, true
						}
					} else {
						return v.Vertex, true
					}
				}
			}
			return nil, false
		}
		return nil, false
	}
}

// ExpressionEvaluator checks the vertex pass the where part of the AST
func (qe Filter) ExpressionEvaluator(expr ast.Expr, variable string, v *vertices.Vertex) bool {
	if inter, ok := expr.(ast.InterpretExpr); ok {
		result := inter.Interpret(variable, v)
		if result, ok := result.(bool); ok {
			return result
		}
	}
	return false
}
