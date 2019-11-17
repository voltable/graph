package expressions

import "reflect"

var _ Expression = (*DefaultExpression)(nil)

type DefaultExpression struct {
}

func (d *DefaultExpression) Compile() Delegate {
	panic("implement me")
}

func (d *DefaultExpression) String() string {
	return ExpressionToString(d)
}

func (d *DefaultExpression) Kind() reflect.Kind {
	return reflect.Invalid
}

func (d *DefaultExpression) Reduce() Expression {
	return baseReduce(d)
}

func (d *DefaultExpression) ReduceAndCheck() Expression {
	return baseReduceAndCheck(d)
}

func (d *DefaultExpression) Accept(visitor ExpressionVisitor) Expression {
	return baseAccept(d, visitor)
}

func (d *DefaultExpression) VisitChildren(visitor ExpressionVisitor) Expression {
	return baseVisitChildren(d, visitor)
}


// Empty expression
func Empty() *DefaultExpression {
	return &DefaultExpression{
	}
}