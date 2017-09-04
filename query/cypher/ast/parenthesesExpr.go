package ast

type Parentheses int

const (
	RPAREN Parentheses = iota // )
	LPAREN                    // (
)

type ParenthesesExpr struct {
	Parentheses
}

var _ Expr = (*ParenthesesExpr)(nil)
var _ InterpretExpr = (*ParenthesesExpr)(nil)

func (ParenthesesExpr) exprNode()      {}
func (ParenthesesExpr) interpretNode() {}

func ParenthesesPrecedence(item ParenthesesExpr) int {
	if item.Parentheses == LPAREN {
		return 11
	} else if item.Parentheses == RPAREN {
		return 12
	} else {
		return 20
	}
}
