package openCypher

import (
	"github.com/voltable/graph/operators"
	"github.com/voltable/graph/operators/ir"
)

type ExecutionPlan struct {
	operators []operators.Operator
}

func NewExecutionPlan() *ExecutionPlan {
	return &ExecutionPlan{
		operators: make([]operators.Operator, 0),
	}
}

func (s *ExecutionPlan) Insert(key int, op operators.Operator, variable ir.Variable) {
	s.operators = append(s.operators, op)
}

func (s *ExecutionPlan) Add(op operators.Operator) {
	s.operators = append(s.operators, op)
}

func (s *ExecutionPlan) Operators() []operators.Operator {
	return s.operators
}
