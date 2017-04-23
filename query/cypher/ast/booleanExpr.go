package ast

// // AndExpr A bitwise or logical AND operation, such as (a And b).
// type AndExpr struct {
// 	BinaryExpr
// }

// // OrExpr A bitwise or logical OR operation, such as (a Or b).
// type OrExpr struct {
// 	BinaryExpr
// }

// // NotExpr A bitwise complement or logical negation operation (Not a).
// type NotExpr struct {
// 	BinaryExpr
// }

// // XorExpr A bitwise or logical XOR operation, such as (a Xor b).
// type XorExpr struct {
// 	BinaryExpr
// }

// func (*AndExpr) exprNode() {}
// func (*OrExpr) exprNode()  {}
// func (*NotExpr) exprNode() {}
// func (*XorExpr) exprNode() {}

type Boolean int

const (
	AND Boolean = iota // AND a bitwise or logical AND operation, such as (a And b).
	OR                 // OR a bitwise or logical OR operation, such as (a Or b).
	NOT                // NOT a bitwise complement or logical negation operation (Not a).
	XOR                // XOR a bitwise or logical XOR operation, such as (a Xor b).
)

type BooleanExpr struct {
	Boolean
	X Expr // left operand
	Y Expr // right operand
}

func (*BooleanExpr) exprNode() {}

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
