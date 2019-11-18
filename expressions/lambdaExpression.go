package expressions

import (
	"github.com/pkg/errors"
	"reflect"
)

var _ Expression = (*LambdaExpression)(nil)

type LambdaExpression struct {
	body Expression
	parameters []*ParameterExpression
}

type Delegate func([]interface{}) interface{}

func (l *LambdaExpression) Compile() Delegate {
	return func(params []interface{}) interface{}{
		builder := NewDynamicExpressionVisitor(params)
		builder.Visit(l)
		return builder.Result()
	}
}

func (l *LambdaExpression) String() string {
	return ExpressionToString(l)
}

func (l *LambdaExpression) Reduce() Expression {
	return baseReduce(l)
}

func (l *LambdaExpression) ReduceAndCheck() Expression {
	return baseReduceAndCheck(l)
}

func (l *LambdaExpression) Accept(visitor ExpressionVisitor) Expression {
	return visitor.VisitLambda(l)
}

func (l *LambdaExpression) VisitChildren(visitor ExpressionVisitor) Expression {
	return baseVisitChildren(l, visitor)
}

func (l *LambdaExpression) Kind() reflect.Kind {
	return reflect.Bool
}

func (l *LambdaExpression) GetBody() Expression {
	return l.body
}

func (l *LambdaExpression) Update(body Expression, parameters []*ParameterExpression) *LambdaExpression {
	if l.body == body {
		equal := true
		for i, parameter := range l.parameters  {
			p := parameters[i]
			if p != parameter {
				equal = false
				break
			}
		}

		if equal {
			return l
		}
	}
	return Lambda(body, parameters...)
}

func Lambda(expr Expression, parameters ...*ParameterExpression) *LambdaExpression {
	if expr == nil {
		panic(errors.Wrapf(ArgumentCannotBeOfTypeVoid, "expr %v", expr))
	}

	return &LambdaExpression{
		body: expr,
		parameters: parameters,
	}
}
