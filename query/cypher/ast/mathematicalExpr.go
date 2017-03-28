package ast

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
