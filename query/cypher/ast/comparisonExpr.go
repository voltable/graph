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
