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

type Parentheses int

const (
	RPAREN Parentheses = iota // )
	LPAREN                    // (
)

type ComparisonExpr struct {
	Comparison
	X Expr // left operand
	Y Expr // right operand
}

type ParenthesesExpr struct {
	Parentheses
}

func (ComparisonExpr) exprNode()  {}
func (ParenthesesExpr) exprNode() {}

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

func ParenthesesPrecedence(item Parentheses) int {
	return 1
}

func ComparisonPrecedence(item ComparisonExpr) int {
	if item.Comparison == EQ {
		return 8
	} else if item.Comparison == NEQ {
		return 8
	} else if item.Comparison == LT {
		return 7
	} else if item.Comparison == LTE {
		return 7
	} else if item.Comparison == GT {
		return 7
	} else if item.Comparison == GTE {
		return 7
	} else if item.Comparison == IS {
		return 7
	} else if item.Comparison == NULL {
		return 7
	} else {
		return 20
	}
}