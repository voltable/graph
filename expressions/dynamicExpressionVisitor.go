package expressions

import "github.com/voltable/graph/expressions/stack"

var _ ExpressionVisitor = (*DynamicExpressionVisitor )(nil)

type DynamicExpressionVisitor  struct {
	stack *stack.Stack
	params []interface{}
}

func NewDynamicExpressionVisitor (params  []interface{}) *DynamicExpressionVisitor {
	return &DynamicExpressionVisitor {
		params:params,
		stack: &stack.Stack{},
	}
}

func (s *DynamicExpressionVisitor ) VisitLambda(expr *LambdaExpression) Expression {
	return baseVisitLambda(s, expr)
}

func (s *DynamicExpressionVisitor ) VisitBinary(expr BinaryExpression) Expression {
	expression := baseVisitBinary(s, expr)
	switch v := expression.(type) {
	case *BinaryArithmeticExpression:
		v.Interpret(s.stack)
	}

	return expression
}

func (s *DynamicExpressionVisitor ) Visit(expr Expression) Expression {
	return baseVisit(s, expr)
}

func (s *DynamicExpressionVisitor ) VisitExtension(expr Expression) Expression {
	return expr.VisitChildren(s)
}

func (s *DynamicExpressionVisitor ) VisitParameter(expr *ParameterExpression) Expression {
	return expr
}

func (s *DynamicExpressionVisitor ) VisitConstant(expr *ConstantExpression) Expression {
	s.stack.Push(expr.GetValue())
	return expr
}

func (s *DynamicExpressionVisitor ) VisitConditional(expr *ConditionalExpression) Expression {
	return baseVisitConditional(s, expr)
}

func (s *DynamicExpressionVisitor) Result() interface{} {

	return s.stack.Pop()
}
