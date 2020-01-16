package openCypher

import (
	"github.com/voltable/graph"
	"github.com/voltable/graph/operators"
	"github.com/voltable/graph/operators/ir"
	"github.com/voltable/graph/widecolumnstore"
)

type QueryPlanner struct {
	ir         StackExpr
	storage    widecolumnstore.Storage
	statistics *graph.Statistics
	plan       *ExecutionPlan
}

func NewQueryPlanner(storage widecolumnstore.Storage, statistics *graph.Statistics, stack StackExpr) *QueryPlanner {
	queryPlanner := &QueryPlanner{}

	var n interface{}
	plan := NewExecutionPlan()
	key := 0

	for stack, n = stack.pop(); n != nil; stack, n = stack.pop() {
		if k, ok := n.(*ir.Return); ok {
			op, _ := operators.NewProduceResults(storage, statistics, k)
			plan.Add(op)

		}
		if k, ok := n.(*ir.Match); ok {
			if k.Pattern != nil {
				for _, part := range k.Pattern.Parts {
					op, _ := operators.NewAllNodesScan(storage, statistics, part.Variable)
					plan.Insert(key, op, part.Variable)
					key++
				}
			}
		}
		if k, ok := n.(*ir.Create); ok {
			if k.Pattern != nil {
				for _, part := range k.Pattern.Parts {
					op, _ := operators.NewCreate(storage, statistics, part.Nodes, part.Relationships)
					plan.Insert(key, op, part.Variable)
					key++
				}
			}
		}
	}

	queryPlanner.plan = plan
	return queryPlanner
}

func (s *QueryPlanner) Execute() *graph.Table {

	results := graph.NewTable()
	//var last operators.Operator
	// s.plan.PreOrderTraverse(func(operator operators.Operator, variable ir.Variable) {
	// 	result = append(result, fmt.Sprintf("%+v %s", operator, variable))

	// })

	var iterator widecolumnstore.Iterator
	for _, op := range s.plan.Operators() {
		if create, ok := op.(*operators.Create); ok {
			iterator = create.Next()
		}
		if allNodesScan, ok := op.(*operators.AllNodesScan); ok {
			iterator = allNodesScan.Next()
		}
		if produceResults, ok := op.(*operators.ProduceResults); ok {
			results = produceResults.Next(iterator)
		}
	}

	return results
}
