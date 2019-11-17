package expressions

import (
	"fmt"
	"github.com/voltable/graph/expressions/interpreter"
	"github.com/voltable/graph/expressions/stack"
	"reflect"
)

var _ BinaryExpression = (*BinaryArithmeticExpression)(nil)

func (BinaryArithmeticExpression) binary() {}

type BinaryArithmeticExpression struct {
	BinaryArithmetic
	kind reflect.Kind
	instruction interpreter.Instruction
	left  Expression // left operand
	right Expression // right operand
}

func (b *BinaryArithmeticExpression) Compile() Delegate {
	panic("implement me")
}

func (b *BinaryArithmeticExpression) String() string {
	return ExpressionToString(b)
}

func (b *BinaryArithmeticExpression) GetConversion() *LambdaExpression {
	return nil
}

func (b *BinaryArithmeticExpression) Update(left TerminalExpression, conversion *LambdaExpression, right TerminalExpression) BinaryExpression {
	return baseUpdate(b, left, conversion, right)
}

func (b *BinaryArithmeticExpression) Reduce() Expression {
	return baseReduce(b)
}

func (b *BinaryArithmeticExpression) ReduceAndCheck() Expression {
	return baseReduceAndCheck(b)
}

func (b *BinaryArithmeticExpression) Accept(visitor ExpressionVisitor) Expression {
	return visitor.VisitBinary(b)
}

func (b *BinaryArithmeticExpression) VisitChildren(visitor ExpressionVisitor) Expression {
	return baseVisitChildren(b, visitor)
}

func (b BinaryArithmeticExpression) Type() Binary {
	return Binary(b.BinaryArithmetic)
}

func (b *BinaryArithmeticExpression) Kind() reflect.Kind {
	return b.kind
}

// GetLeft return value store in left side
func (b *BinaryArithmeticExpression) GetLeft() Expression {
	return b.left
}

// GetRight return value store in right side
func (b *BinaryArithmeticExpression) GetRight() Expression {
	return b.right
}


func (b *BinaryArithmeticExpression)Interpret(stack *stack.Stack) error {
	return b.instruction.Run(stack)
}

func Add(left TerminalExpression, right TerminalExpression) *BinaryArithmeticExpression {
	instruction, err := interpreter.NewAddInstruction(left.GetValue())
	if err != nil {
		panic(err)
	}

	if left.Kind() == right.Kind() {
		return &BinaryArithmeticExpression{
			BinaryArithmetic: add,
			instruction:instruction,
			left:             left,
			right:            right,
		}
	}

	panic(fmt.Sprintf("mismatch of types %T %T", left, right))
}

func Divide(left TerminalExpression, right TerminalExpression) *BinaryArithmeticExpression {
	instruction, err := interpreter.NewDivideInstruction(left.GetValue())
	if err != nil {
		panic(err)
	}

	if left.Kind() == right.Kind() {
		return &BinaryArithmeticExpression{
			BinaryArithmetic: divide,
			instruction:instruction,
			left:             left,
			right:            right,
		}
	}

	panic(fmt.Sprintf("mismatch of types %T %T", left, right))
}

func Modulo(left TerminalExpression, right TerminalExpression) *BinaryArithmeticExpression {
	instruction, err := interpreter.NewModuloInstruction(left.GetValue())
	if err != nil {
		panic(err)
	}

	if left.Kind() == right.Kind() {
		return &BinaryArithmeticExpression{
			BinaryArithmetic: modulo,
			instruction:instruction,
			left:             left,
			right:            right,
		}
	}

	panic(fmt.Sprintf("mismatch of types %T %T", left, right))
}

func Multiply(left TerminalExpression, right TerminalExpression) *BinaryArithmeticExpression {
	instruction, err := interpreter.NewMultiplyInstruction(left.GetValue())
	if err != nil {
		panic(err)
	}

	if left.Kind() == right.Kind() {
		return &BinaryArithmeticExpression{
			BinaryArithmetic: multiply,
			instruction:instruction,
			left:             left,
			right:            right,
		}
	}

	panic(fmt.Sprintf("mismatch of types %T %T", left, right))
}

func Power(left TerminalExpression, right TerminalExpression) *BinaryArithmeticExpression {
	instruction, err := interpreter.NewPowerInstruction(left.GetValue())
	if err != nil {
		panic(err)
	}

	if left.Kind() == right.Kind() {
		return &BinaryArithmeticExpression{
			BinaryArithmetic: power,
			instruction:instruction,
			left:             left,
			right:            right,
		}
	}

	panic(fmt.Sprintf("mismatch of types %T %T", left, right))
}

func Subtract(left TerminalExpression, right TerminalExpression) *BinaryArithmeticExpression {
	instruction, err := interpreter.NewSubtractInstruction(left.GetValue())
	if err != nil {
		panic(err)
	}

	if left.Kind() == right.Kind() {
		return &BinaryArithmeticExpression{
			BinaryArithmetic: subtract,
			instruction:      instruction,
			left:             left,
			right:            right,
		}
	}

	panic(fmt.Sprintf("mismatch of types %T %T", left, right))
}