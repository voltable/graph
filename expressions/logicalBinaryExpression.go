package expressions

import (
	"reflect"
)

var _ BinaryExpression = (*LogicalBinaryExpression)(nil)

type LogicalBinaryExpression struct {
	Logical
	left  Expression // left operand
	right Expression // right operand
}

func (*LogicalBinaryExpression) binary() {}

func (e *LogicalBinaryExpression) String() string {
	return ExpressionToString(e)
}

func (e *LogicalBinaryExpression) Reduce() Expression {
	return e
}

func (e *LogicalBinaryExpression) ReduceAndCheck() Expression {
	return baseReduceAndCheck(e)
}

func (e *LogicalBinaryExpression) Accept(visitor ExpressionVisitor) Expression {
	return visitor.VisitBinary(e)
}

func (e *LogicalBinaryExpression) VisitChildren(visitor ExpressionVisitor) Expression {
	return baseVisitChildren(e, visitor)
}

func (e *LogicalBinaryExpression) Kind() reflect.Kind {
	return reflect.Bool
}

func (e *LogicalBinaryExpression) GetLeft() Expression {
	return e.left
}

func (e *LogicalBinaryExpression) GetRight() Expression {
	return e.right
}

func (e *LogicalBinaryExpression) Type() Binary {
	return Binary(e.Logical)
}

func (e *LogicalBinaryExpression) Update(left, right TerminalExpression) BinaryExpression {
	return baseUpdate(e, left, right)
}

func Equal(left, right Expression) *LogicalBinaryExpression {
	if left == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	if right == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &LogicalBinaryExpression{
		Logical: equal,
		left:    left,
		right:   right,
	}
}

func NotEqual(left, right Expression) *LogicalBinaryExpression {
	if left == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}
	if right == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &LogicalBinaryExpression{
		Logical: notEqual,
		left:    left,
		right:   right,
	}
}

func LessThan(left, right Expression) *LogicalBinaryExpression {
	if left == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	if right == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &LogicalBinaryExpression{
		Logical: lessThan,
		left:    left,
		right:   right,
	}
}

func LessThanOrEqual(left, right Expression) *LogicalBinaryExpression {
	if left == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	if right == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &LogicalBinaryExpression{
		Logical: lessThanOrEqual,
		left:    left,
		right:   right,
	}
}

func GreaterThan(left, right Expression) *LogicalBinaryExpression {
	if left == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	if right == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &LogicalBinaryExpression{
		Logical: greaterThan,
		left:    left,
		right:   right,
	}
}

func GreaterThanOrEqual(left, right Expression) *LogicalBinaryExpression {
	if left == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	if right == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &LogicalBinaryExpression{
		Logical: greaterThanOrEqual,
		left:    left,
		right:   right,
	}
}

func IsNil(left, right Expression) *LogicalBinaryExpression {
	if left == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	if right == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &LogicalBinaryExpression{
		Logical: isNil,
		left:    left,
		right:   right,
	}
}

func IsNotNil(left, right Expression) *LogicalBinaryExpression {
	if left == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	if right == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &LogicalBinaryExpression{
		Logical: isNotNil,
		left:    left,
		right:   right,
	}
}
