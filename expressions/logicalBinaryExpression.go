package expressions

import (
	"reflect"
)


var _ BinaryExpression = (*LogicalBinaryExpression)(nil)

func (LogicalBinaryExpression) binary() {}

type LogicalBinaryExpression struct {
	Logical
	left  Expression // left operand
	right Expression // right operand
}

func (e LogicalBinaryExpression) String() string {
	return ExpressionToString(e)
}

func (e LogicalBinaryExpression) Reduce() (Expression, error) {
	return e, nil
}

func (e LogicalBinaryExpression) ReduceAndCheck() (Expression, error) {
	return baseReduceAndCheck(e)
}

func (e LogicalBinaryExpression) Accept(visitor ExpressionVisitor) (Expression, error) {
	return visitor.VisitBinary(e)
}

func (e LogicalBinaryExpression) VisitChildren(visitor ExpressionVisitor) (Expression, error) {
	return baseVisitChildren(e, visitor)
}

func (e LogicalBinaryExpression) Kind() reflect.Kind {
	return reflect.Bool
}

func (e LogicalBinaryExpression) GetLeft() Expression {
	return e.left
}

func (e LogicalBinaryExpression) GetRight() Expression {
	return  e.right
}

func (e LogicalBinaryExpression) Type() Binary {
	return Binary(e.Logical)
}

func (e LogicalBinaryExpression) Update(left TerminalExpression, conversion *LambdaExpression, right TerminalExpression) (BinaryExpression, error) {
	return baseUpdate(e, left, conversion, right)
}

func (e LogicalBinaryExpression) GetConversion() *LambdaExpression {
	return nil
}


func Equal(left, right TerminalExpression) (*LogicalBinaryExpression, error) {
	if left == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	if right == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	return &LogicalBinaryExpression{
		Logical: equal,
		left: left,
		right: right,
	}, nil
}

func NotEqual(left, right TerminalExpression) (*LogicalBinaryExpression, error) {
	if left == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	if right == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	return &LogicalBinaryExpression{
		Logical: notEqual,
		left: left,
		right: right,
	}, nil
}

func LessThan(left, right TerminalExpression) (*LogicalBinaryExpression, error) {
	if left == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	if right == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	return &LogicalBinaryExpression{
		Logical: lessThan,
		left: left,
		right: right,
	}, nil
}

func LessThanOrEqual(left, right TerminalExpression) (*LogicalBinaryExpression, error) {
	if left == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	if right == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	return &LogicalBinaryExpression{
		Logical: lessThanOrEqual,
		left: left,
		right: right,
	}, nil
}

func GreaterThan(left, right TerminalExpression) (*LogicalBinaryExpression, error) {
	if left == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	if right == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	return &LogicalBinaryExpression{
		Logical: greaterThan,
		left: left,
		right: right,
	}, nil
}

func GreaterThanOrEqual(left, right TerminalExpression) (*LogicalBinaryExpression, error) {
	if left == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	if right == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	return &LogicalBinaryExpression{
		Logical: greaterThanOrEqual,
		left: left,
		right: right,
	}, nil
}

func IsNil(left, right TerminalExpression) (*LogicalBinaryExpression, error) {
	if left == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	if right == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	return &LogicalBinaryExpression{
		Logical: isNil,
		left: left,
		right: right,
	}, nil
}

func IsNotNil(left, right TerminalExpression) (*LogicalBinaryExpression, error) {
	if left == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	if right == nil {
		return nil, ArgumentCannotBeOfTypeVoid
	}

	return &LogicalBinaryExpression{
		Logical: isNotNil,
		left: left,
		right: right,
	}, nil
}