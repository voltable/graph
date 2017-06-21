package ast

// Comparison operators
type Comparison int

const (
	// EQ equality
	EQ Comparison = iota // =
	// NEQ inequality
	NEQ // <>
	// LT less than
	LT // <
	// LTE less than or equal to
	LTE // <=
	// GT greater than
	GT // >
	// GTE greater than or equal to
	GTE // >=
	// IS used of IS NULL
	IS // IS
	// NULL used for IS NULL
	NULL // NULL
)

// ComparisonExpr comparison expression
type ComparisonExpr struct {
	Comparison
	X Expr // left operand
	Y Expr // right operand
}

func (ComparisonExpr) exprNode() {}

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
