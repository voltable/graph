package parser

import (
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
)

type stackExpr []ast.Expr

func (s stackExpr) Push(v ast.Expr) stackExpr {
	return append(s, v)
}

func (s stackExpr) Pop() (stackExpr, ast.Expr, bool) {
	l := len(s)
	if l > 0 {
		//return s[:l-1], s[l-1], true
		return s[1:], s[0], true
	}
	return s, nil, false
}
