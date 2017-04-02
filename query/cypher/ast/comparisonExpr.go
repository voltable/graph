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
	BinaryExpr
	Comparison
}

func (*ComparisonExpr) exprNode() {}
