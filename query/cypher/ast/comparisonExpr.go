package ast

type Comparison int

const (
	EQ   Comparison = iota // =
	NEQ                    // <>
	LT                     // <
	LTE                    // <=
	GT                     // >
	GTE                    // >=
	IS                     // IS
	NULL                   // NULL
)

type ComparisonExpr struct {
	Comparison
	X Expr // left operand
	Y Expr // right operand
}

func (*ComparisonExpr) exprNode() {}

func (b *ComparisonExpr) GetX() Expr {
	return b.X
}

func (b *ComparisonExpr) GetY() Expr {
	return b.Y
}

func (b *ComparisonExpr) SetX(x Expr) {
	b.X = x
}

func (b *ComparisonExpr) SetY(y Expr) {
	b.Y = y
}
