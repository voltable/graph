package parser

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

type StackExpr []ast.Expr

func (s StackExpr) Push(v ast.Expr) StackExpr {
	return append(s, v)
}

func (s StackExpr) Pop() (StackExpr, ast.Expr, bool) {
	l := len(s)
	if l > 0 {
		return s[:l-1], s[l-1], true
		//return s[1:], s[0], true
	}
	return s, nil, false
}

// UpdateStack builds up the AST by Shunting the stack if a expression supports children
func (s StackExpr) UpdateStack(expr ast.Expr) StackExpr {
	var item ast.Expr
	var temp ast.Expr
	s = s.Push(expr)

	for len(s) > 1 {
		s, item, _ = s.Pop()

		compar, ok := item.(ast.OperatorExpr)
		if !ok {
			old := item
			s, item, _ = s.Pop()
			s = s.Push(old)
		}

		if ok {
			if compar.GetX() == nil {
				s, temp, _ = s.Pop()
				compar.SetX(temp)
				s = s.Push(item)
				if ast.IsOperatorWithFreeXorY(temp) {
					s.Push(temp)
				}
				return s
			} else if compar.GetY() == nil {
				s, temp, _ = s.Pop()
				compar.SetY(temp)
				s = s.Push(item)
				if ast.IsOperatorWithFreeXorY(temp) {
					s.Push(temp)
				}
				return s
			}
		} else {
			s = s.Push(item)
		}
	}

	return s
}

func (s StackExpr) setCompar(item ast.Expr, temp ast.Expr, set func(ast.Expr)) StackExpr {
	set(temp)
	s = s.Push(item)
	if ast.IsOperatorWithFreeXorY(temp) {
		s.Push(temp)
	}

	return s
}
