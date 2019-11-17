package expressions

import "reflect"

var _ TerminalExpression = (*ConstantExpression)(nil)

// ConstantExpression represents a constant expression.
type ConstantExpression struct {
	data interface{} // denoted object; or nil
	kind reflect.Kind
}

func (s *ConstantExpression) Compile() Delegate {
	panic("implement me")
}

func (s *ConstantExpression) Kind() reflect.Kind {
	return s.kind
}

func (s *ConstantExpression) Reduce() Expression {
	return baseReduce(s)
}

func (s *ConstantExpression) ReduceAndCheck() Expression {
	return baseReduceAndCheck(s)
}

func (s *ConstantExpression) Accept(visitor ExpressionVisitor) Expression {
	return visitor.VisitConstant(s)
}

func (s *ConstantExpression) VisitChildren(visitor ExpressionVisitor) Expression {
	return baseVisitChildren(s, visitor)
}

func (s *ConstantExpression) GetValue() interface{} {
	return s.data
}

func (s *ConstantExpression) String() string {
	return ExpressionToString(s)
}

func Constant(i interface{}) *ConstantExpression{
	if i == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &ConstantExpression{
		data: i,
		kind: reflect.TypeOf(i).Kind(),
	}
}