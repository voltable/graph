package expressions

import "reflect"

var _ Expression = (*InvocationExpression)(nil)
var _ ArgumentProvider = (*InvocationExpression)(nil)

type InvocationExpression struct {
	expression Expression
	arguments  []Expression
}

// Invoke creates a InvocationExpression that applies a lambda expression to one argument expression.
func Invoke(expression Expression, arguments ...Expression) *InvocationExpression {
	return &InvocationExpression{
		expression: expression,
		arguments:  arguments,
	}
}

func (l *InvocationExpression) Accept(visitor ExpressionVisitor) Expression {
	return visitor.VisitInvocation(l)
}

func (l *InvocationExpression) Kind() reflect.Kind {
	return reflect.Func
}

func (l *InvocationExpression) Reduce() Expression {
	return baseReduce(l)
}

func (l *InvocationExpression) ReduceAndCheck() Expression {
	return baseReduceAndCheck(l)
}

func (l *InvocationExpression) VisitChildren(visitor ExpressionVisitor) Expression {
	return baseVisitChildren(l, visitor)
}

func (l *InvocationExpression) String() string {
	return ExpressionToString(l)
}

func (l *InvocationExpression) Expression() Expression {
	return l.expression
}

func (l *InvocationExpression) Arguments() []Expression {
	return l.arguments
}

func (l *InvocationExpression) Rewrite(expr Expression, arguments []Expression) *InvocationExpression {
	return Invoke(expr, arguments...)
}
