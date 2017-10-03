package ast

// A BinaryExpr node represents a binary expression.
type BinaryExpr struct {
	Left  Expr // left operand
	Right Expr // right operand
}

func (BinaryExpr) exprNode() {}
