package ast

type ParenthesesExpr struct {
	Parentheses
}

var _ Expr = (*ParenthesesExpr)(nil)
var _ InterpretExpr = (*ParenthesesExpr)(nil)

func (ParenthesesExpr) exprNode()      {}
func (ParenthesesExpr) interpretNode() {}

func (p *ParenthesesExpr) Interpret(variable string, prop Interpret) interface{} {
	return nil
}

func ParenthesesPrecedence(item ParenthesesExpr) int {
	if item.Parentheses == LPAREN {
		return 11
	} else if item.Parentheses == RPAREN {
		return 12
	} else {
		return 20
	}
}
