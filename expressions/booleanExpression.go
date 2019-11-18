package expressions

import "reflect"

var _ BinaryExpression = (*BooleanExpression)(nil)

func (BooleanExpression) binary() {}

// BooleanExpression boolean expression
type BooleanExpression struct {
	Boolean
	kind reflect.Kind
	left  Expression // left operand
	right Expression // right operand
}

func (b *BooleanExpression) String() string {
	return ExpressionToString(b)
}

func (b *BooleanExpression) Reduce() Expression {
	return baseReduce(b)
}

func (b *BooleanExpression) ReduceAndCheck() Expression {
	return baseReduceAndCheck(b)
}

func (b *BooleanExpression) Accept(visitor ExpressionVisitor) Expression {
	return visitor.VisitBinary(b)
}

func (b *BooleanExpression) VisitChildren(visitor ExpressionVisitor) Expression {
	return baseVisitChildren(b, visitor)
}

func (b *BooleanExpression) Kind() reflect.Kind {
	return b.kind
}

func (b *BooleanExpression) Type() Binary {
	return Binary(b.Boolean)
}

func (b *BooleanExpression) Update(left TerminalExpression, conversion *LambdaExpression, right TerminalExpression) BinaryExpression {
	return baseUpdate(b, left, conversion, right)
}

func (b *BooleanExpression) GetConversion() *LambdaExpression {
	return nil
}

// NewBooleanExpr creates a BooleanExpr
func NewBooleanExpr(boolean Boolean, left Expression, right Expression) *BooleanExpression {
	return &BooleanExpression{
		Boolean: boolean,
		left: left,
		right: right}
}

// GetLeft return value store in left side
func (b *BooleanExpression) GetLeft() Expression {
	return b.left
}

// GetRight return value store in right side
func (b *BooleanExpression) GetRight() Expression {
	return b.right
}

// SetLeft stores the Expr in left side
func (b *BooleanExpression) SetLeft(left Expression) {
	b.left = left
}

// SetRight stores the Expr in right side
func (b *BooleanExpression) SetRight(right Expression) {
	b.right = right
}


// BooleanPrecedence returns the precedence (order of importance)
func BooleanPrecedence(item BooleanExpression) int {
	if item.Boolean == and {
		return 9
	} else if item.Boolean == or {
		return 11
	} else if item.Boolean == xor {
		return 10
	}
	return 20
}

func And(left, right TerminalExpression) *BooleanExpression {
	if left == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	if right == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &BooleanExpression{
		Boolean: and,
		left: left,
		right: right,
	}
}

func Or(left, right TerminalExpression) *BooleanExpression {
	if left == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	if right == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &BooleanExpression{
		Boolean: or,
		left: left,
		right: right,
	}
}


func Xor(left, right TerminalExpression) *BooleanExpression {
	if left == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	if right == nil {
		panic(ArgumentCannotBeOfTypeVoid)
	}

	return &BooleanExpression{
		Boolean: xor,
		left: left,
		right: right,
	}
}