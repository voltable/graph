package expressions

import "reflect"

var _ TerminalExpression = (*ParameterExpression)(nil)

// A ParameterExpression node represents a parameter expression.
type ParameterExpression  struct {
	name string
	kind reflect.Kind
}

func (p *ParameterExpression) GetValue() interface{} {
	return p.kind
}

func (p *ParameterExpression) Compile() Delegate {
	panic("implement me")
}

func (p *ParameterExpression) String() string {
	return ExpressionToString(p)
}

func (p *ParameterExpression) Kind() reflect.Kind {
	return p.kind
}


func (p *ParameterExpression) ReduceAndCheck() Expression {
	return baseReduceAndCheck(p)
}

func (p *ParameterExpression) VisitChildren(visitor ExpressionVisitor) Expression {
	return baseVisitChildren(p, visitor)
}

func (p *ParameterExpression) Reduce() Expression {
	return baseReduce(p)
}

func (p *ParameterExpression) Accept(visitor ExpressionVisitor) Expression {
	return visitor.VisitParameter(p)
}

func (p *ParameterExpression) GetName() string {
	return p.name
}

func Parameter(kind reflect.Kind, name string) *ParameterExpression {
	return &ParameterExpression{
		kind: kind,
		name: name,
	}
}

func Variable(kind reflect.Kind, name string) *ParameterExpression {
	return &ParameterExpression{
		kind: kind,
		name: name,
	}
}