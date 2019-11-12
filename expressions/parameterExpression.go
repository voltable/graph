package expressions

import "reflect"

var _ Expression = (*ParameterExpression)(nil)

// A ParameterExpression node represents a parameter expression.
type ParameterExpression  struct {
	name string
	kind reflect.Kind
}

func (p *ParameterExpression) String() string {
	return ExpressionToString(p)
}

func (p *ParameterExpression) Kind() reflect.Kind {
	return p.kind
}


func (p *ParameterExpression) ReduceAndCheck() (Expression, error) {
	return baseReduceAndCheck(p)
}

func (p *ParameterExpression) VisitChildren(visitor ExpressionVisitor) (Expression, error) {
	return baseVisitChildren(p, visitor)
}

func (p *ParameterExpression) Reduce() (Expression, error) {
	return baseReduce(p)
}

func (p *ParameterExpression) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitParameter(p)
}

func (p *ParameterExpression) GetName() string {
	return p.name
}

func Parameter(i interface{}, name string) (*ParameterExpression, error) {
	if i == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}
	return &ParameterExpression{
		kind: reflect.TypeOf(i).Kind(),
		name: name,
	}, nil
}

func Variable(i interface{}, name string) (*ParameterExpression, error) {
	if i == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	return &ParameterExpression{
		kind: reflect.TypeOf(i).Kind(),
		name: name,
	}, nil
}