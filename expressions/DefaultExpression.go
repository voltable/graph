package expressions

import "reflect"

type DefaultExpression struct {
}

var _ Expression = (*DefaultExpression)(nil)

func (d *DefaultExpression) String() string {
	return ExpressionToString(d)
}

func (d *DefaultExpression) Kind() reflect.Kind {
	return reflect.Invalid
}

func (d *DefaultExpression) Reduce() (Expression, error) {
	return baseReduce(d)
}

func (d *DefaultExpression) ReduceAndCheck() (Expression, error) {
	return baseReduceAndCheck(d)
}

func (d *DefaultExpression) Accept(visitor ExpressionVisitor) (Expression, error) {
	return baseAccept(d, visitor)
}

func (d *DefaultExpression) VisitChildren(visitor ExpressionVisitor) (Expression, error) {
	return baseVisitChildren(d, visitor)
}



// Empty expression
func Empty() *DefaultExpression {
	return &DefaultExpression{
	}
}