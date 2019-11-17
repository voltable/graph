package expressions

import "reflect"

// ConstantExpression represents a constant expression.
type ConstantExpression struct {
	data interface{} // denoted object; or nil
	kind reflect.Kind
}

func (s *ConstantExpression) Compile() func() {
	panic("implement me")
}

var _ TerminalExpression = (*ConstantExpression)(nil)

func (s *ConstantExpression) Kind() reflect.Kind {
	return s.kind
}

func (s *ConstantExpression) Reduce() (Expression, error) {
	return baseReduce(s)
}

func (s *ConstantExpression) ReduceAndCheck() (Expression, error) {
	return baseReduceAndCheck(s)
}

func (s *ConstantExpression) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitConstant(s)
}

func (s *ConstantExpression) VisitChildren(visitor ExpressionVisitor) (Expression, error) {
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