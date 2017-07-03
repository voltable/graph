package ast

type Boolean int

const (
	AND Boolean = iota // AND a bitwise or logical AND operation, such as (a And b).
	OR                 // OR a bitwise or logical OR operation, such as (a Or b).
	//NOT                // NOT a bitwise complement or logical negation operation (Not a).
	XOR // XOR a bitwise or logical XOR operation, such as (a Xor b).
)

type BooleanExpr struct {
	Boolean
	X Expr // left operand
	Y Expr // right operand
}

// NotExpr a bitwise complement or logical negation operation (Not a).
type NotExpr struct {
	X Expr // left operand
}

func (BooleanExpr) exprNode() {}
func (NotExpr) exprNode()     {}

func (b *BooleanExpr) GetX() Expr {
	return b.X
}

func (b *BooleanExpr) GetY() Expr {
	return b.Y
}

func (b *BooleanExpr) SetX(x Expr) {
	b.X = x
}

func (b *BooleanExpr) SetY(y Expr) {
	b.Y = y
}

func (b *NotExpr) SetX(x Expr) {
	b.X = x
}

func BooleanPrecedence(item BooleanExpr) int {
	if item.Boolean == AND {
		return 9
	} else if item.Boolean == OR {
		return 11
	} else if item.Boolean == XOR {
		return 10
	} else {
		return 20
	}
}

func NotPrecedence(item NotExpr) int {
	return 13
}
