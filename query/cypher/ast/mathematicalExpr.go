package ast

// MathematicalExpr all maths nodes implement the MathematicalExpr interface.
type MathematicalExpr interface {
	mathNode()
}

// AddExpr An addition operation, such as a + b, for numeric operands.
type AddExpr struct {
	BinaryExpr
}

// MultiplyExpr A multiplication operation, such as (a * b), for numeric operands.
type MultiplyExpr struct {
	BinaryExpr
}

// DivideExpr A division operation, such as (a / b), for numeric operands.
type DivideExpr struct {
	BinaryExpr
}

// SubtractExpr A subtraction operation, such as (a - b), for numeric operands.
type SubtractExpr struct {
	BinaryExpr
}

// ModuloExpr An arithmetic remainder operation, such as (a % b).
type ModuloExpr struct {
	BinaryExpr
}

// PowerExpr A mathematical operation that raises a number to a power, such as (a ^ b) in Visual Basic.
type PowerExpr struct {
	BinaryExpr
}

func (*AddExpr) exprNode()      {}
func (*MultiplyExpr) exprNode() {}
func (*DivideExpr) exprNode()   {}
func (*SubtractExpr) exprNode() {}
func (*ModuloExpr) exprNode()   {}
func (*PowerExpr) exprNode()    {}

func (*AddExpr) mathNode()      {}
func (*MultiplyExpr) mathNode() {}
func (*DivideExpr) mathNode()   {}
func (*SubtractExpr) mathNode() {}
func (*ModuloExpr) mathNode()   {}
func (*PowerExpr) mathNode()    {}

// MathPrecedence returns the precedence (order of importance)
func MathPrecedence(item MathematicalExpr) int {
	if _, ok := item.(*AddExpr); ok {
		return 2
	} else if _, ok := item.(*SubtractExpr); ok {
		return 2
	} else if _, ok := item.(*DivideExpr); ok {
		return 3
	} else if _, ok := item.(*MultiplyExpr); ok {
		return 3
	} else if _, ok := item.(*ModuloExpr); ok {
		return 4
	} else if _, ok := item.(*PowerExpr); ok {
		return 4
	} else {
		return 20
	}
}
