package expressions

import (
	"reflect"
)

type LambdaExpression struct {

}

func (l LambdaExpression) String() string {
	panic("implement me")
}

func (l LambdaExpression) Reduce() (Expression, error) {
	panic("implement me")
}

func (l LambdaExpression) ReduceAndCheck() (Expression, error) {
	panic("implement me")
}

func (l LambdaExpression) Accept(visitor ExpressionVisitor) (Expression, error) {
	panic("implement me")
}

func (l LambdaExpression) VisitChildren(visitor ExpressionVisitor) (Expression, error) {
	panic("implement me")
}

func (l LambdaExpression) Kind() reflect.Kind {
	panic("implement me")
}

var _ Expression = (*LambdaExpression)(nil)
