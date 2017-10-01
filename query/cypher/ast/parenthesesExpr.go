package ast

import "github.com/RossMerr/Caudex.Graph/vertices"

type ParenthesesExpr struct {
	Parentheses
}

var _ Expr = (*ParenthesesExpr)(nil)
var _ InterpretExpr = (*ParenthesesExpr)(nil)

func (ParenthesesExpr) exprNode()      {}
func (ParenthesesExpr) interpretNode() {}

func (p *ParenthesesExpr) Interpret(vertex *vertices.Vertex) interface{} {
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
