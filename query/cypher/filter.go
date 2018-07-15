package cypher

import (
	"github.com/RossMerr/Caudex.Graph/keyvalue"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

// CypherFilter the interface of the filter struct
type CypherFilter interface {
	Filter(query.IteratorFrontier, ast.Expr) query.IteratorFrontier
}

// Filter is use to filter the traveral results over the where expression in the AST
type Filter struct {
}

// NewFilter create a new filter
func NewFilter() *Filter {
	return &Filter{}
}

// Filter a IteratorFrontier so that all results pass the Where Expression in the AST
func (qe Filter) Filter(i query.IteratorFrontier, predicate ast.Expr) query.IteratorFrontier {
	check := true
	return func() (*query.Frontier, bool) {
		frontier, ok := i()
		for ok {
			check = true
			if frontier.Len() > 0 {
				optimalPath := frontier.OptimalPath()
				frontier.Clear()
				if len(optimalPath) == 0 {
					check = false
				}
				if predicate != nil {
					for _, i := range optimalPath {
						if !qe.ExpressionEvaluator(predicate, i.Variable, i.Object) {
							check = false
							break
						}
					}
				}
			} else {
				check = false
			}

			if check {
				return frontier, ok
			}

			frontier, ok = i()
		}

		return frontier, ok
	}
}

// ExpressionEvaluator checks the vertex pass the where part of the AST
func (qe Filter) ExpressionEvaluator(expr ast.Expr, variable string, prop *keyvalue.KeyValue) bool {
	if inter, ok := expr.(ast.InterpretExpr); ok {
		result := inter.Interpret(variable, prop)
		if result, ok := result.(bool); ok {
			return result
		}
	}
	return false
}
