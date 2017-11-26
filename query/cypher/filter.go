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
	length := 0
	position := 0
	frontier, ok := i()
	var queue []interface{}

	return func() (interface{}, bool) {
		for ok != query.Failed {
			if position == 0 {
				if frontier == nil {
					return nil, false
				}
				if frontier.Len() > 0 {
					queue = frontier.OptimalPath()
					length = len(queue)
				}
			}
			if position < length {
				i := queue[position]
				position++
				if v, ok := i.(*query.FrontierVertex); ok {

					if predicate != nil {
						if qe.ExpressionEvaluator(predicate, v.Variable, v.Vertex) {
							return v.Vertex, true
						}
					} else {
						return v.Vertex, true
					}
				} else if e, ok := i.(*query.FrontierEdge); ok {
					// todo need to run predicate over edge
					return e.Edge, true
				}
			}

			frontier, ok = i()
			position = 0
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
